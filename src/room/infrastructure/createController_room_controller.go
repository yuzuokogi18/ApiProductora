package infrastructure

import (
	"database/sql"
	"productor/src/room/application"
	"productor/src/room/domain"
	"github.com/rabbitmq/amqp091-go"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateRoomController struct {
	useCase *application.CreateRoomUseCase
}

func NewCreateRoomController(useCase *application.CreateRoomUseCase) *CreateRoomController {
	return &CreateRoomController{useCase: useCase}
}

func (controller *CreateRoomController) Execute(c *gin.Context) {
	var room domain.Room

	if err := c.BindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos no válidos"})
		return
	}

	if err := controller.useCase.Run(&room); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la habitación"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Habitación creada y enviada a RabbitMQ"})
}

func InitializeRoomUseCase(db *sql.DB, rabbitmqCh *amqp091.Channel) *application.CreateRoomUseCase {
	roomPgRepository := NewRoomPgRepository(db)
	roomRabbitmqRepository := NewRoomRabbitmqRepository(rabbitmqCh)

	return application.NewCreateRoomUseCase(roomRabbitmqRepository, roomPgRepository)
}
