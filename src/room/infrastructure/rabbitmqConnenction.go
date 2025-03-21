package infrastructure

import (
	"fmt"
	"log"
	amqp "github.com/rabbitmq/amqp091-go"
	"productor/src/room/domain"
)

// RoomRabbitmqRepository es el repositorio que maneja la comunicación con RabbitMQ para las habitaciones
type RoomRabbitmqRepository struct {
	ch *amqp.Channel
}

// NewRoomRabbitmqRepository crea un nuevo repositorio para interactuar con RabbitMQ
func NewRoomRabbitmqRepository(ch *amqp.Channel) *RoomRabbitmqRepository {
	// Declarar el exchange 'logs' para las habitaciones
	if err := ch.ExchangeDeclare(
		"logs",   // Nombre del exchange (mantenemos 'logs' como en la implementación original)
		"fanout", // Tipo del exchange (fanout envía a todas las colas)
		true,     // Durable (el exchange persiste incluso si RabbitMQ se reinicia)
		false,    // Auto-deleted
		false,    // Internal
		false,    // No-wait
		nil,      // Argumentos
	); err != nil {
		log.Fatalf("Error al declarar el exchange de habitaciones: %v", err)
	}

	return &RoomRabbitmqRepository{ch: ch}
}

// Save envía un mensaje con los datos de la habitación a RabbitMQ
func (repo *RoomRabbitmqRepository) Save(room *domain.Room) error {
	// Crear el mensaje que se enviará a RabbitMQ
	message := fmt.Sprintf("Nueva habitación en el hotel %d, Tipo: %s, Capacidad: %d, Precio: %.2f",
		room.HotelId, room.Type, room.Capacity, room.Price)

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
