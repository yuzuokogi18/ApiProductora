package domain

type IHotelPg interface {
	Save(Hotel *Hotel) error
	GetAll() ([]Hotel, error)
}
