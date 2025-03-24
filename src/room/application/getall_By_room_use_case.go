package application

import (
	"productor/src/room/domain"
	"fmt"
)

type GetAllRoomsByHotelUseCase struct {
	roomPgRepository domain.IRoomPg
	roomRabbitmqRepository domain.IRoomRabbitqm
}

func NewGetAllRoomsByHotelUseCase(roomRabbitmqRepository domain.IRoomRabbitqm, roomPgRepository domain.IRoomPg) *GetAllRoomsByHotelUseCase {
	return &GetAllRoomsByHotelUseCase{roomRabbitmqRepository: roomRabbitmqRepository, roomPgRepository: roomPgRepository}
}

func (usecase *GetAllRoomsByHotelUseCase) Run(hotelID int) ([]domain.Room, error) {
	rooms, err := usecase.roomPgRepository.GetRoomsByHotel(hotelID)
	if err != nil {
		return nil, fmt.Errorf("Error al obtener las habitaciones: %v", err)
	}

	errSendMessage := usecase.roomRabbitmqRepository.SendAllRooms(rooms)
	if errSendMessage != nil {
		return nil, fmt.Errorf("Error al enviar las habitaciones a RabbitMQ: %v", errSendMessage)
	}

	return rooms, nil
}
