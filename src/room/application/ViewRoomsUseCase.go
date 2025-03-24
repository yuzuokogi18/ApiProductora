package application

import (
	"productor/src/room/domain"
	"log"
)

type ViewRoomUseCase struct {
	pgRepository       domain.IRoomPg
	rabbitqmRepository domain.IRoomRabbitqm
}

func NewViewRoomUseCase(pgRepository domain.IRoomPg, rabbitqmRepository domain.IRoomRabbitqm) *ViewRoomUseCase {
	return &ViewRoomUseCase{pgRepository: pgRepository, rabbitqmRepository: rabbitqmRepository}
}

func (usecase *ViewRoomUseCase) SetRoom(pgRepository domain.IRoomPg, rabbitqmRepository domain.IRoomRabbitqm) {
	usecase.pgRepository = pgRepository
	usecase.rabbitqmRepository = rabbitqmRepository
}

func (usecase *ViewRoomUseCase) Run(roomID string) (*domain.Room, error) {
	room, err := usecase.pgRepository.FindByID(roomID)
	if err != nil {
		log.Printf("Error al obtener la habitación de la base de datos: %v", err)
		return nil, err
	}

	errSendMessage := usecase.rabbitqmRepository.SendViewRequest(roomID)
	if errSendMessage != nil {
		log.Printf("Error al enviar la solicitud de visualización a RabbitMQ: %v", errSendMessage)
		return nil, errSendMessage
	}

	return room, nil
}
