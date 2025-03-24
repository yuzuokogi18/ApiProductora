package application

import (
	"productor/src/hotel/domain"
	"log"
)

type GetByIdHotelUseCase struct {
	pgRepository       domain.IHotelPg
	rabbitqmRepository domain.IHotelRabbitqm
}

func NewGetByIdHotelUseCase(pgRepository domain.IHotelPg, rabbitqmRepository domain.IHotelRabbitqm) *GetByIdHotelUseCase {
	return &GetByIdHotelUseCase{
		pgRepository:       pgRepository,
		rabbitqmRepository: rabbitqmRepository,
	}
}

func (usecase *GetByIdHotelUseCase) SetHotel(pgRepository domain.IHotelPg, rabbitqmRepository domain.IHotelRabbitqm) {
	usecase.pgRepository = pgRepository
	usecase.rabbitqmRepository = rabbitqmRepository
}

func (usecase *GetByIdHotelUseCase) Run(id int) (*domain.Hotel, error) {
	hotel, err := usecase.pgRepository.GetById(id)
	if err != nil {
		log.Printf("Error al obtener el hotel con ID %d: %v", id, err)
		return nil, err
	}

	
	errSendMessage := usecase.rabbitqmRepository.Save(hotel)
	if errSendMessage != nil {
		log.Printf("Error al enviar el hotel a RabbitMQ: %v", errSendMessage)
		return nil, errSendMessage
	}

	return hotel, nil
}
