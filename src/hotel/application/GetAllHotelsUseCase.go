package application

import (
	"productor/src/hotel/domain"
	"log"
)

type GetAllHotelsUseCase struct {
	pgRepository       domain.IHotelPg
	rabbitqmRepository domain.IHotelRabbitqm
}

func NewGetAllHotelsUseCase(pgRepository domain.IHotelPg, rabbitqmRepository domain.IHotelRabbitqm) *GetAllHotelsUseCase {
	return &GetAllHotelsUseCase{
		pgRepository:       pgRepository,
		rabbitqmRepository: rabbitqmRepository,
	}
}

func (usecase *GetAllHotelsUseCase) SetHotel(pgRepository domain.IHotelPg, rabbitqmRepository domain.IHotelRabbitqm) {
	usecase.pgRepository = pgRepository
	usecase.rabbitqmRepository = rabbitqmRepository
}

func (usecase *GetAllHotelsUseCase) Run() ([]domain.Hotel, error) {
	// Obtener todos los hoteles de la base de datos
	hotels, err := usecase.pgRepository.GetAll()
	if err != nil {
		log.Printf("Error al obtener los hoteles: %v", err)
		return nil, err
	}

	
	errSendMessage := usecase.rabbitqmRepository.SaveAll(hotels) 
	if errSendMessage != nil {
		log.Printf("Error al enviar los hoteles a RabbitMQ: %v", errSendMessage)
		return nil, errSendMessage
	}

	return hotels, nil
}
