package domain

type Reservation struct {
	Id          int32  `json:"id"`
	CustomerName string `json:"customer_name"`
	RoomType    string `json:"room_type"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Price       float32 `json:"price"`
}

func NewReservation(customerName, roomType, startDate, endDate string, price float32) *Reservation {
	return &Reservation{CustomerName: customerName, RoomType: roomType, StartDate: startDate, EndDate: endDate, Price: price}
}
  