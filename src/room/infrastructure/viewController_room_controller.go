package infrastructure

import (
	"database/sql"
	"productor/src/room/application"
	"github.com/rabbitmq/amqp091-go"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewRoomController struct {
	useCase *application.ViewRoomUseCase
}

func NewViewRoomController(useCase *application.ViewRoomUseCase) *ViewRoomController {
	return &ViewRoomController{useCase: useCase}
}

func (controller *ViewRoomController) Execute(c *gin.Context) {
	roomID := c.Param("id")

	room, err := controller.useCase.Run(roomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener la habitación"})
		return
	}

	c.JSON(http.StatusOK, room)
}

func InitializeViewRoomUseCase(db *sql.DB, rabbitmqCh *amqp091.Channel) *application.ViewRoomUseCase {
	roomPgRepository := NewRoomPgRepository(db)
	roomRabbitmqRepository := NewRoomRabbitmqRepository(rabbitmqCh)

	return application.NewViewRoomUseCase(roomPgRepository, roomRabbitmqRepository)
}
