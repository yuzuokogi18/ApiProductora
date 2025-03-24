package infrastructure

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go" 
	"net/http"
	"strconv" 
	"productor/src/room/application"
)

type GetAllRoomsByHotelController struct {
	useCase *application.GetAllRoomsByHotelUseCase
}

func NewGetAllRoomsByHotelController(useCase *application.GetAllRoomsByHotelUseCase) *GetAllRoomsByHotelController {
	return &GetAllRoomsByHotelController{useCase: useCase}
}

func (controller *GetAllRoomsByHotelController) Execute(c *gin.Context) {
	hotelID := c.Param("hotelID")
	if hotelID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El ID del hotel es requerido"})
		return
	}

	hotelIDInt, err := strconv.Atoi(hotelID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de hotel no válido"})
		return
	}

	rooms, err := controller.useCase.Run(hotelIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las habitaciones"})
		return
	}

	
	c.JSON(http.StatusOK, rooms)
}

func InitializeGetAllRoomsByHotelUseCase(db *sql.DB, rabbitmqCh *amqp091.Channel) *application.GetAllRoomsByHotelUseCase {
	roomPgRepository := NewRoomPgRepository(db)
	roomRabbitmqRepository := NewRoomRabbitmqRepository(rabbitmqCh)

	return application.NewGetAllRoomsByHotelUseCase(roomRabbitmqRepository, roomPgRepository)
}
