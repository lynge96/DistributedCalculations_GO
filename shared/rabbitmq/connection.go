package rabbitmq

import (
	"fmt"
	"log/slog"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Connection struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   string
}

func NewConnection(connString string, queue string) (*Connection, error) {
	conn, err := dial(connString, 10)

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

	slog.Info("queue declared:", "queue", q.Name)

	return &Connection{
		Conn:    conn,
		Channel: ch,
		Queue:   q.Name,
	}, nil
}

func dial(connString string, maxRetries int) (*amqp.Connection, error) {
	var err error
	for i := range maxRetries {
		conn, err := amqp.Dial(connString)
		if err == nil {
			return conn, nil
		}
		slog.Info("failed to connect to RabbitMQ, retrying in 5s...", "attempt", i+1, "maxRetries", maxRetries)
		time.Sleep(5 * time.Second)
	}
	return nil, fmt.Errorf("failed to connect to RabbitMQ after %d retries: %w", maxRetries, err)
}

func (c *Connection) Close() {
	_ = c.Channel.Close()
	_ = c.Conn.Close()
}
