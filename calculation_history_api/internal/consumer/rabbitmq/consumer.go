package rabbitmq

import (
	"encoding/json"
	"log"
	"shared/models"

	amqp "github.com/rabbitmq/amqp091-go"
)

type HistoryStore interface {
	Add(entry models.CalculationResult)
}

type Consumer struct {
	store   HistoryStore
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   string
}

func NewConsumer(store HistoryStore, connString string, queue string) (*Consumer, error) {

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
		return nil, err
	}
	log.Printf("Queue declared: %s", q.Name)

	return &Consumer{
		store:   store,
		conn:    conn,
		channel: ch,
		queue:   q.Name,
	}, nil
}

func (c *Consumer) Start() error {

	msgs, err := c.channel.Consume(
		c.queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	log.Printf("Consumer started, waiting for messages on queue: %s", c.queue)

	for msg := range msgs {
		var result models.CalculationResult
		if err := json.Unmarshal(msg.Body, &result); err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			continue
		}

		c.store.Add(result)
		log.Printf("Received and stored result: %+v", result)
	}

	return nil
}
