package domain

type IRoomRabbitqm interface {
	Save(room *Room) error
	SendViewRequest(roomID string) error
	SendAllRooms(rooms []Room) error
}