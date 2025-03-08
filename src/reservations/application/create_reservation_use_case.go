package application

import (
	"productor/src/reservations/domain"
	"log"
)

type CreateReservationUseCase struct {
	rabbitqmRepository domain.IReservationRabbitqm
	pgRepository       domain.IReservationPg
}

func NewCreateReservationUseCase(rabbitqmRepository domain.IReservationRabbitqm, pgRepository domain.IReservationPg) *CreateReservationUseCase {
	return &CreateReservationUseCase{rabbitqmRepository: rabbitqmRepository, pgRepository: pgRepository}
}

func (usecase *CreateReservationUseCase) SetReservation(pgRepository domain.IReservationPg, rabbitqmRepository domain.IReservationRabbitqm) {
	usecase.pgRepository = pgRepository
	usecase.rabbitqmRepository = rabbitqmRepository
}

func (usecase *CreateReservationUseCase) Run(reservation *domain.Reservation) error {
	// Guardar la reservación en la base de datos
	err := usecase.pgRepository.Save(reservation)
	if err != nil {
		log.Printf("Error al guardar la reservación en la base de datos: %v", err)
		return err
	}

	// Enviar el mensaje a RabbitMQ
	errSendMessage := usecase.rabbitqmRepository.Save(reservation)
	if errSendMessage != nil {
		log.Printf("Error al enviar el mensaje a RabbitMQ: %v", errSendMessage)
		return errSendMessage
	}

	return nil
}
