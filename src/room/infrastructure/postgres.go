package infrastructure

import (
	"productor/src/room/domain"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type RoomPgRepository struct {
	db *sql.DB
}

func NewRoomPgRepository(db *sql.DB) *RoomPgRepository {
	return &RoomPgRepository{db: db}
}

func (repo *RoomPgRepository) Save(room *domain.Room) error {
	query := "INSERT INTO rooms (hotel_id, type, capacity, price) VALUES ($1, $2, $3, $4)"
	_, err := repo.db.Exec(query, room.HotelId, room.Type, room.Capacity, room.Price)
	if err != nil {
		return fmt.Errorf("Error al guardar la habitación: %v", err)
	}

	return nil
}
