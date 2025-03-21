package domain

type IRoomRabbitqm interface {
	Save(Room *Room) error
}
