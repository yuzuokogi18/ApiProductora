package domain

type Room struct {
	Id       int32   `json:"id"`
	HotelId  int32   `json:"hotel_id"`
	Type     string  `json:"type"`
	Capacity int     `json:"capacity"`
	Price    float32 `json:"price"`
}

func NewRoom(hotelId int32, roomType string, capacity int, price float32) *Room {
	return &Room{HotelId: hotelId, Type: roomType, Capacity: capacity, Price: price}
}
