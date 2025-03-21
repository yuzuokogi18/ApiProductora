package domain

type IHotelRabbitqm interface {
	Save(Hotel *Hotel) error
	SaveAll(hotels []Hotel) error
}
