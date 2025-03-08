package infrastructure

import (
	"productor/src/reservations/application"
	"productor/src/reservations/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateReservationController struct {
	useCase *application.CreateReservationUseCase
}

func NewCreateReservationController(useCase *application.CreateReservationUseCase) *CreateReservationController {
	return &CreateReservationController{useCase: useCase}
}

func (controller *CreateReservationController) Execute(c *gin.Context) {
	var reservation domain.Reservation
	
	// Bind JSON to reservation object
	if err := c.BindJSON(&reservation); err != nil {
		// If JSON binding fails, return a Bad Request error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos no válidos"})
		return
	}

	// Call the use case to process the reservation creation logic
	if err := controller.useCase.Run(&reservation); err != nil {
		// If the use case fails, return an Internal Server Error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la reservación"})
		return
	}

	// If everything goes well, return a success response
	c.JSON(http.StatusOK, gin.H{"message": "Reservación creada y enviada a RabbitMQ"})
}
