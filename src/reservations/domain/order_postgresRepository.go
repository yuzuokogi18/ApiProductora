package domain

type IReservationPg interface {
	Save(Reservation *Reservation) error
}
