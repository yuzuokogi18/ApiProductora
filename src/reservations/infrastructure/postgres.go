package infrastructure

import (
	"productor/src/reservations/domain"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PgRepository struct {
	db *sql.DB
}

func NewPgRepository(db *sql.DB) *PgRepository {
	return &PgRepository{db: db}
}

func (repo *PgRepository) Save(reservation *domain.Reservation) error {
	query := "INSERT INTO reservations (customer_name, room_type, start_date, end_date, price) VALUES ($1, $2, $3, $4, $5)"
	_, err := repo.db.Exec(query, reservation.CustomerName, reservation.RoomType, reservation.StartDate, reservation.EndDate, reservation.Price)
	if err != nil {
		return fmt.Errorf("Error al guardar la reservación: %v", err)
	}

	return nil
}
