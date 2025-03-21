package application

import (
	"productor/src/hotel/domain"
	"log"
)

type CreateHotelUseCase struct {
	rabbitqmRepository domain.IHotelRabbitqm 
	pgRepository       domain.IHotelPg
}

func NewCreateHotelUseCase(rabbitqmRepository domain.IHotelRabbitqm , pgRepository domain.IHotelPg) *CreateHotelUseCase {
	return &CreateHotelUseCase{rabbitqmRepository: rabbitqmRepository, pgRepository: pgRepository}
}

func (usecase *CreateHotelUseCase) SetHotel(pgRepository domain.IHotelPg, rabbitqmRepository domain.IHotelRabbitqm ) {
	usecase.pgRepository = pgRepository
	usecase.rabbitqmRepository = rabbitqmRepository
}

func (usecase *CreateHotelUseCase) Run(hotel *domain.Hotel) error {

	err := usecase.pgRepository.Save(hotel)
	if err != nil {
		log.Printf("Error al guardar el hotel en la base de datos: %v", err)
		return err
	}

	errSendMessage := usecase.rabbitqmRepository.Save(hotel)
	if errSendMessage != nil {
		log.Printf("Error al enviar el mensaje a RabbitMQ: %v", errSendMessage)
		return errSendMessage
	}

	return nil
}
