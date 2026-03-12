package rabbitmq

import (
	"encoding/json"
	"log/slog"
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
		slog.Error("failed to create RabbitMQ connection", "error", err, "queue", queue, "connString", connString)
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
	slog.Info("consumer started, waiting for messages on queue", "queue", c.conn.Queue)

	for msg := range msgs {
		var result models.CalculationResult
		if err := json.Unmarshal(msg.Body, &result); err != nil {
			slog.Warn("failed to unmarshal message", "error", err, "body", string(msg.Body))
			continue
		}

		c.store.Add(result)
		slog.Info("received and stored result", "result", result)
	}

	return nil
}

func (p *Consumer) Close() {
	p.conn.Close()
}
