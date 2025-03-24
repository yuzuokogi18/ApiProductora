package domain

type IRoomPg interface {
	Save(room *Room) error
	FindByID(roomID string) (*Room, error)
	GetRoomsByHotel(hotelID int) ([]Room, error)
}