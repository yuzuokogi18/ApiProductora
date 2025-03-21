package domain

type Hotel struct {
	Id       int32   `json:"id"`
	Name     string  `json:"name"`
	Location string  `json:"location"`
	Stars    int     `json:"stars"`
	Price    float32 `json:"price"`
}

func NewHotel(name, location string, stars int, price float32) *Hotel {
	return &Hotel{Name: name, Location: location, Stars: stars, Price: price}
}
