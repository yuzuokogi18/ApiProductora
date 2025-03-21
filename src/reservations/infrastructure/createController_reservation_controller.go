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
	
	
	if err := c.BindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos no válidos"})
		return
	}

	if err := controller.useCase.Run(&reservation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la reservación"})
		return
	}


	c.JSON(http.StatusOK, gin.H{"message": "Reservación creada y enviada a RabbitMQ"})
}
