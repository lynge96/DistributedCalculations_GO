package rabbitmq

import (
	"encoding/json"
	"log"
	"shared/models"
	"shared/rabbitmq"
)

type HistoryStore interface {
	Add(entry models.CalculationResult)
}

type Consumer struct {
	store HistoryStore
	conn  *rabbitmq.Connection
}

func NewConsumer(store HistoryStore, connString string, queue string) (*Consumer, error) {
	conn, err := rabbitmq.NewConnection(connString, queue)
	if err != nil {
		return nil, err
	}
	return &Consumer{store: store, conn: conn}, nil
}

func (c *Consumer) Start() error {

	msgs, err := c.conn.Channel.Consume(
		c.conn.Queue,
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
	log.Printf("Consumer started, waiting for messages on queue: %s", c.conn.Queue)

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
