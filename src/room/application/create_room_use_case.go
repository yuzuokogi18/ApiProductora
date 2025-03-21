package application
import (
	"productor/src/room/domain"
	"log"
)

type CreateRoomUseCase struct {
	rabbitqmRepository domain.IRoomRabbitqm
	pgRepository       domain.IRoomPg
}

func NewCreateRoomUseCase(rabbitqmRepository domain.IRoomRabbitqm, pgRepository domain.IRoomPg) *CreateRoomUseCase {
	return &CreateRoomUseCase{rabbitqmRepository: rabbitqmRepository, pgRepository: pgRepository}
}

func (usecase *CreateRoomUseCase) SetRoom(pgRepository domain.IRoomPg, rabbitqmRepository domain.IRoomRabbitqm) {
	usecase.pgRepository = pgRepository
	usecase.rabbitqmRepository = rabbitqmRepository
}

func (usecase *CreateRoomUseCase) Run(room *domain.Room) error {

	err := usecase.pgRepository.Save(room)
	if err != nil {
		log.Printf("Error al guardar la habitación en la base de datos: %v", err)
		return err
	}


	errSendMessage := usecase.rabbitqmRepository.Save(room)
	if errSendMessage != nil {
		log.Printf("Error al enviar el mensaje a RabbitMQ: %v", errSendMessage)
		return errSendMessage
	}

	return nil
}
