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

func (repo *RoomPgRepository) FindByID(roomID string) (*domain.Room, error) {
	query := "SELECT id, hotel_id, type, capacity, price FROM rooms WHERE id = $1"
	row := repo.db.QueryRow(query, roomID)

	var room domain.Room
	if err := row.Scan(&room.Id, &room.HotelId, &room.Type, &room.Capacity, &room.Price); err != nil {
		return nil, fmt.Errorf("Error al obtener la habitación: %v", err)
	}

	return &room, nil
}

func (repo *RoomPgRepository) GetRoomsByHotel(hotelID int) ([]domain.Room, error) {
	query := "SELECT id, hotel_id, type, capacity, price FROM rooms WHERE hotel_id = $1"
	rows, err := repo.db.Query(query, hotelID)
	if err != nil {
		return nil, fmt.Errorf("Error al obtener las habitaciones del hotel con ID %d: %v", hotelID, err)
	}
	defer rows.Close()

	var rooms []domain.Room
	for rows.Next() {
		var room domain.Room
		if err := rows.Scan(&room.Id, &room.HotelId, &room.Type, &room.Capacity, &room.Price); err != nil {
			return nil, fmt.Errorf("Error al leer los datos de la habitación: %v", err)
		}
		rooms = append(rooms, room)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error durante la iteración de filas: %v", err)
	}

	return rooms, nil
}

func (repo *RoomPgRepository) UpdateAvailability(roomID string, available bool) error {
	query := "UPDATE rooms SET available = $1 WHERE id = $2"
	_, err := repo.db.Exec(query, available, roomID)
	if err != nil {
		return fmt.Errorf("Error al actualizar la disponibilidad de la habitación: %v", err)
	}

	return nil
}
