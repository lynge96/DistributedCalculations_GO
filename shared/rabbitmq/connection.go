package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Connection struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   string
}

func NewConnection(connString string, queue string) (*Connection, error) {
	conn, err := amqp.Dial(connString)
	if err != nil {
		return nil, err
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
