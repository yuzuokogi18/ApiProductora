package domain

type IReservationRabbitqm interface {
	Save(Reservation *Reservation) error
}
