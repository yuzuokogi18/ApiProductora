package infrastructure

import (
	"database/sql"
	"productor/src/hotel/application"
	"github.com/rabbitmq/amqp091-go"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetByIdHotelController struct {
	useCase *application.GetByIdHotelUseCase
}

func NewGetByIdHotelController(useCase *application.GetByIdHotelUseCase) *GetByIdHotelController {
	return &GetByIdHotelController{useCase: useCase}
}

func (controller *GetByIdHotelController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	hotel, err := controller.useCase.Run(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el hotel"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"hotel": hotel})
}

func InitializeGetByIdHotelUseCase(db *sql.DB, rabbitmqCh *amqp091.Channel) *application.GetByIdHotelUseCase {
	hotelPgRepository := NewHotelPgRepository(db)
	hotelRabbitmqRepository := NewHotelRabbitmqRepository(rabbitmqCh)

	return application.NewGetByIdHotelUseCase(hotelPgRepository, hotelRabbitmqRepository)
}
