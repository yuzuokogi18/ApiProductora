package infrastructure

import (
	"productor/src/hotel/domain"
	"database/sql"
	"fmt"
)

type HotelPgRepository struct {
	db *sql.DB
}

// Nueva instancia de HotelPgRepository
func NewHotelPgRepository(db *sql.DB) *HotelPgRepository {
	return &HotelPgRepository{db: db}
}

// El método Save guarda un hotel en la base de datos
func (repo *HotelPgRepository) Save(hotel *domain.Hotel) error {
	query := "INSERT INTO hotels (name, location, stars, price) VALUES ($1, $2, $3, $4)"
	_, err := repo.db.Exec(query, hotel.Name, hotel.Location, hotel.Stars, hotel.Price)
	if err != nil {
		return fmt.Errorf("Error al guardar el hotel: %v", err)
	}
	return nil
}

// El método GetAll obtiene todos los hoteles de la base de datos
func (repo *HotelPgRepository) GetAll() ([]domain.Hotel, error) {
	query := "SELECT id, name, location, stars, price FROM hotels"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar la consulta: %v", err)
	}
	defer rows.Close()

	var hotels []domain.Hotel
	for rows.Next() {
		var hotel domain.Hotel
		if err := rows.Scan(&hotel.Id, &hotel.Name, &hotel.Location, &hotel.Stars, &hotel.Price); err != nil {
			return nil, fmt.Errorf("Error al escanear los datos: %v", err)
		}
		hotels = append(hotels, hotel)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error al iterar sobre las filas: %v", err)
	}

	return hotels, nil
}

// El método GetById obtiene un hotel por su ID
func (repo *HotelPgRepository) GetById(id int) (*domain.Hotel, error) {
	query := "SELECT id, name, location, stars, price FROM hotels WHERE id = $1"
	row := repo.db.QueryRow(query, id)

	var hotel domain.Hotel
	if err := row.Scan(&hotel.Id, &hotel.Name, &hotel.Location, &hotel.Stars, &hotel.Price); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("No se encontró el hotel con ID %d", id)
		}
		return nil, fmt.Errorf("Error al obtener el hotel: %v", err)
	}

	return &hotel, nil
}
