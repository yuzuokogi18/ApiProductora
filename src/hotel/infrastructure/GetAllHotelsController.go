package infrastructure

import (
	"database/sql"
	"productor/src/hotel/application"
	"github.com/rabbitmq/amqp091-go"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllHotelsController struct {
	useCase *application.GetAllHotelsUseCase
}

func NewGetAllHotelsController(useCase *application.GetAllHotelsUseCase) *GetAllHotelsController {
	return &GetAllHotelsController{useCase: useCase}
}

func (controller *GetAllHotelsController) Execute(c *gin.Context) {
	
	hotels, err := controller.useCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los hoteles"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"hotels": hotels})
}


func InitializeGetAllHotelsUseCase(db *sql.DB, rabbitmqCh *amqp091.Channel) *application.GetAllHotelsUseCase {
	hotelPgRepository := NewHotelPgRepository(db)
	hotelRabbitmqRepository := NewHotelRabbitmqRepository(rabbitmqCh)

	return application.NewGetAllHotelsUseCase(hotelPgRepository, hotelRabbitmqRepository)
}
