package rabbitmq

import (
	"encoding/json"
	"log/slog"
	"shared/models"
	"shared/rabbitmq"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	conn *rabbitmq.Connection
}

func NewPublisher(connString string, queue string) (*Publisher, error) {

	conn, err := rabbitmq.NewConnection(connString, queue)
	if err != nil {
		slog.Error("failed to create RabbitMQ connection:", "error", err)
		return nil, err
	}
	return &Publisher{conn: conn}, nil
}

func (p *Publisher) Publish(message models.CalculationResult) error {

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	slog.Info("publishing message:", "body", string(body))
	return p.conn.Channel.Publish(
		"",
		p.conn.Queue,
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

func (p *Publisher) Close() {
	p.conn.Close()
}
