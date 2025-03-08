package core

import (
	"log"
	"productor/src/reservations/application"
	"productor/src/reservations/infrastructure"
	"github.com/gin-gonic/gin"
	"productor/src/reservations/domain"
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

    pgRepository := infrastructure.NewPgRepository(pgConn.DB)
	rabbitqmRepository := infrastructure.NewRabbitRepository(rabbitmqCh.ch)

	createReservationUseCase := application.NewCreateReservationUseCase(rabbitqmRepository, pgRepository)

	router := gin.Default()

	router.POST("/reservation", func(c *gin.Context) {
		var reservation domain.Reservation
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

	log.Fatal(router.Run(":8080"))
}
