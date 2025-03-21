package domain

type IRoomPg interface {
	Save(Room *Room) error
}
