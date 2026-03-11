package main

import (
	"calculator_api/internal/messaging/rabbitmq"
	"log"
	"net/http"
	"shared/configuration"

	"calculator_api/internal/api"
	"calculator_api/internal/calculator"
)

func main() {

	connString := configuration.GetEnv("RABBITMQ_URL", "amqp://guest:guest@raspberrypi:5672/")
	queue := configuration.GetEnv("RABBITMQ_QUEUE", "calculations")
	port := configuration.GetEnv("PORT", "8080")

	publisher, err := rabbitmq.NewPublisher(connString, queue)
	if err != nil {
		log.Fatalf("Failed to create publisher: %v", err)
	}
	defer publisher.Close()

	parser := &calculator.GovaluateParser{}
	service := calculator.NewService(parser, publisher)

	handler := api.NewHandler(service)
	router := api.NewRouter(handler)

	log.Printf("Server running on port :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
