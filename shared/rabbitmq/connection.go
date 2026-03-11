package rabbitmq

import (
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Connection struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   string
}

func NewConnection(connString string, queue string) (*Connection, error) {
	var conn *amqp.Connection
	var err error

	maxRetries := 10
	for i := range maxRetries {
		conn, err = amqp.Dial(connString)
		if err == nil {
			break
		}
		log.Printf("Failed to connect to RabbitMQ, retrying in 5s... (%d/%d)", i+1, maxRetries)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ after %d retries: %w", maxRetries, err)
	}

	ch, err := conn.Channel()
	if err != nil {
		_ = conn.Close()
		return nil, err
	}

	q, err := ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		_ = conn.Close()
		return nil, err
	}

	log.Printf("Queue declared: %s", q.Name)

	return &Connection{
		Conn:    conn,
		Channel: ch,
		Queue:   q.Name,
	}, nil
}

func (c *Connection) Close() {
	_ = c.Channel.Close()
	_ = c.Conn.Close()
}
