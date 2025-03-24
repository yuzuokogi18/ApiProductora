package core

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	reservationApp "productor/src/reservations/application"
	reservationInfra "productor/src/reservations/infrastructure"
	hotelApp "productor/src/hotel/application"
	hotelInfra "productor/src/hotel/infrastructure"
	roomApp "productor/src/room/application"
	roomInfra "productor/src/room/infrastructure"
	reservationDomain "productor/src/reservations/domain"
	hotelDomain "productor/src/hotel/domain"
	roomDomain "productor/src/room/domain"
)

func IniciarRutas() {
	pgConn, err := GetDBPool()
	if err != nil {
		log.Fatalf("Error al obtener la conexión a la base de datos: %v", err)
	}

	rabbitmqCh, err := GetChannel()
	if err != nil {
		log.Fatalf("Error al obtener la conexión a RabbitMQ: %v", err)
	}

	pgRepository := reservationInfra.NewPgRepository(pgConn.DB)
	rabbitmqRepository := reservationInfra.NewRabbitRepository(rabbitmqCh.ch)

	hotelRepository := hotelInfra.NewHotelPgRepository(pgConn.DB)
	hotelRabbitmqRepository := hotelInfra.NewHotelRabbitmqRepository(rabbitmqCh.ch)

	createHotelUseCase := hotelApp.NewCreateHotelUseCase(hotelRabbitmqRepository, hotelRepository)
	getHotelUseCase := hotelApp.NewGetAllHotelsUseCase(hotelRepository, hotelRabbitmqRepository)
	getHotelByIdUseCase := hotelApp.NewGetByIdHotelUseCase(hotelRepository, hotelRabbitmqRepository)

	createReservationUseCase := reservationApp.NewCreateReservationUseCase(rabbitmqRepository, pgRepository)

	roomPgRepository := roomInfra.NewRoomPgRepository(pgConn.DB)
	roomRabbitmqRepository := roomInfra.NewRoomRabbitmqRepository(rabbitmqCh.ch)

	createRoomUseCase := roomApp.NewCreateRoomUseCase(roomRabbitmqRepository, roomPgRepository)
	getAllRoomsByHotelUseCase := roomApp.NewGetAllRoomsByHotelUseCase(roomRabbitmqRepository, roomPgRepository)

	router := gin.Default()

	// Configuración CORS antes de las rutas
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	router.POST("/reservation", func(c *gin.Context) {
		var reservation reservationDomain.Reservation
		if err := c.ShouldBindJSON(&reservation); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		err := createReservationUseCase.Run(&reservation)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error al crear la reservación"})
			return
		}

		c.JSON(201, gin.H{"message": "Reservación creada exitosamente"})
	})


	router.POST("/hotel", func(c *gin.Context) {
		var hotel hotelDomain.Hotel
		if err := c.ShouldBindJSON(&hotel); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		err := createHotelUseCase.Run(&hotel)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error al crear el hotel"})
			return
		}

		c.JSON(201, gin.H{"message": "Hotel creado exitosamente"})
	})

	
	router.GET("/hotel", func(c *gin.Context) {
		hotels, err := getHotelUseCase.Run()
		if err != nil {
			c.JSON(500, gin.H{"error": "Error al obtener los hoteles"})
			return
		}

		c.JSON(200, gin.H{"hotels": hotels})
	})
	router.GET("/hotel/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(400, gin.H{"error": "ID de hotel inválido"})
			return
		}

		hotel, err := getHotelByIdUseCase.Run(id)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"hotel": hotel})
	})

	router.POST("/room", func(c *gin.Context) {
		var room roomDomain.Room
		if err := c.ShouldBindJSON(&room); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		err := createRoomUseCase.Run(&room)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error al crear la habitación"})
			return
		}

		c.JSON(201, gin.H{"message": "Habitación creada exitosamente"})
	})
	router.GET("/hotel/:id/room", func(c *gin.Context) {
		hotelID := c.Param("id")
		hotelIDInt, err := strconv.Atoi(hotelID)
		if err != nil {
			c.JSON(400, gin.H{"error": "ID de hotel inválido"})
			return
		}
	
		rooms, err := getAllRoomsByHotelUseCase.Run(hotelIDInt)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error al obtener las habitaciones"})
			return
		}
	
		c.JSON(200, rooms)
	})
	

	router.GET("/room/:id", func(c *gin.Context) {
		roomID := c.Param("id")
		room, err := roomPgRepository.FindByID(roomID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error al obtener la habitación"})
			return
		}
		c.JSON(200, gin.H{"room": room})
	})

	log.Fatal(router.Run(":8080"))
}
