module calculator_api

go 1.26.1

require (
	github.com/Knetic/govaluate v3.0.0+incompatible
	github.com/google/uuid v1.6.0
	shared v0.0.0
)

require github.com/rabbitmq/amqp091-go v1.10.0

replace shared => ../shared