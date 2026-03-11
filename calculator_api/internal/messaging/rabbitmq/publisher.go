package rabbitmq

import (
	"encoding/json"
	"log"
	"shared/models"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   string
}

func NewPublisher(connString string, queue string) (*Publisher, error) {

	conn, err := amqp.Dial(connString)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
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
		return nil, err
	}
	log.Printf("Queue declared: %s", q.Name)

	return &Publisher{
		conn:    conn,
		channel: ch,
		queue:   q.Name,
	}, nil
}

func (p *Publisher) Publish(message models.CalculationResult) error {

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return p.channel.Publish(
		"",
		p.queue,
		false,
		false,
		amqp.Publishing{
			ContentType:     "application/json",
			ContentEncoding: "utf-8",
			DeliveryMode:    amqp.Persistent,
			Body:            body,
			Type:            "calculation-result",
		},
	)
}
