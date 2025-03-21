package infrastructure

import (
	"encoding/json"
	"fmt"
	"log"
	amqp "github.com/rabbitmq/amqp091-go"
	"productor/src/hotel/domain"
)

// HotelRabbitmqRepository es el repositorio que maneja la comunicación con RabbitMQ para los hoteles
type HotelRabbitmqRepository struct {
	ch *amqp.Channel
}

// NewHotelRabbitmqRepository crea un nuevo repositorio para interactuar con RabbitMQ
func NewHotelRabbitmqRepository(ch *amqp.Channel) *HotelRabbitmqRepository {
	// Declarar el exchange 'logs' para los hoteles
	if err := ch.ExchangeDeclare(
		"logs",   // Nombre del exchange (mantenemos 'logs' como en la implementación original)
		"fanout", // Tipo del exchange (fanout envía a todas las colas)
		true,     // Durable (el exchange persiste incluso si RabbitMQ se reinicia)
		false,    // Auto-deleted
		false,    // Internal
		false,    // No-wait
		nil,      // Argumentos
	); err != nil {
		log.Fatalf("Error al declarar el exchange de hoteles: %v", err)
	}

	return &HotelRabbitmqRepository{ch: ch}
}

// Save envía un mensaje con los datos del hotel a RabbitMQ
func (repo *HotelRabbitmqRepository) Save(hotel *domain.Hotel) error {
	// Crear el mensaje que se enviará a RabbitMQ
	message := fmt.Sprintf("Nuevo hotel: %s, Ubicación: %s, Estrellas: %d, Precio: %.2f",
		hotel.Name, hotel.Location, hotel.Stars, hotel.Price)

	// Publicar el mensaje en el exchange 'logs'
	err := repo.ch.Publish(
		"logs",           // Exchange
		"",               // Routing key (vacío para fanout)
		false,            // Mandatory
		false,            // Immediate
		amqp.Publishing{
			ContentType: "text/plain", // Tipo de contenido
			Body:        []byte(message), // Cuerpo del mensaje (en texto plano)
		},
	)

	if err != nil {
		log.Printf("Error al enviar el mensaje a RabbitMQ: %v", err)
		return err
	}

	log.Printf(" [x] Enviado: %s", message)
	return nil
}

// SaveAll envía una lista de hoteles como mensaje único a RabbitMQ
func (repo *HotelRabbitmqRepository) SaveAll(hotels []domain.Hotel) error {
	// Serializar la lista de hoteles a formato JSON
	body, err := json.Marshal(hotels)
	if err != nil {
		log.Printf("Error al serializar los hoteles: %v", err)
		return err
	}

	// Publicar el mensaje en el exchange 'logs' con los hoteles serializados
	err = repo.ch.Publish(
		"logs",           // Exchange
		"",               // Routing key (vacío para fanout)
		false,            // Mandatory
		false,            // Immediate
		amqp.Publishing{
			ContentType: "application/json", // Tipo de contenido
			Body:        body, // Cuerpo del mensaje (en formato JSON)
		},
	)

	if err != nil {
		log.Printf("Error al enviar los hoteles a RabbitMQ: %v", err)
		return err
	}

	log.Printf(" [x] Enviado lista de hoteles")
	return nil
}
