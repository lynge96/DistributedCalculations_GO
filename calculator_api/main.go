package main

import (
	"calculator_api/internal/messaging/rabbitmq"
	"log"
	"net/http"

	"calculator_api/internal/api"
	"calculator_api/internal/calculator"
)

func main() {

	publisher, err := rabbitmq.NewPublisher("amqp://guest:guest@raspberrypi:5672/", "calculations")
	if err != nil {
		log.Fatal(err)
	}

	parser := &calculator.GovaluateParser{}
	service := calculator.NewService(parser, publisher)

	handler := api.NewHandler(service)
	router := api.NewRouter(handler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
