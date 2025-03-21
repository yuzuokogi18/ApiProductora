package infrastructure

import (
	"database/sql"
	"productor/src/hotel/application"
	"productor/src/hotel/domain"
	"github.com/rabbitmq/amqp091-go"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateHotelController struct {
	useCase *application.CreateHotelUseCase
}

func NewCreateHotelController(useCase *application.CreateHotelUseCase) *CreateHotelController {
	return &CreateHotelController{useCase: useCase}
}

func (controller *CreateHotelController) Execute(c *gin.Context) {
	var hotel domain.Hotel

	if err := c.BindJSON(&hotel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos no válidos"})
		return
	}

	if err := controller.useCase.Run(&hotel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el hotel"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hotel creado y enviado a RabbitMQ"})
}

func InitializeHotelUseCase(db *sql.DB, rabbitmqCh *amqp091.Channel) *application.CreateHotelUseCase {
	hotelPgRepository := NewHotelPgRepository(db)
	hotelRabbitmqRepository := NewHotelRabbitmqRepository(rabbitmqCh)

	return application.NewCreateHotelUseCase(hotelRabbitmqRepository, hotelPgRepository)
}
