package domain

type IHotelPg interface {
	Save(Hotel *Hotel) error
	GetAll() ([]Hotel, error)
	GetById(id int) (*Hotel, error)
}
