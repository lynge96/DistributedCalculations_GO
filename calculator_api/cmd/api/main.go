package main

import (
	"calculator_api/internal/messaging/rabbitmq"
	"log"
	"net/http"

	"calculator_api/internal/api"
	"calculator_api/internal/calculator"
)

func main() {

	parser := &calculator.GovaluateParser{}
	service := calculator.NewService(parser)

	publisher, err := rabbitmq.NewPublisher("amqp://guest:guest@raspberrypi:5672/", "calculations")
	if err != nil {
		log.Fatal(err)
	}

	handler := api.NewHandler(service, publisher)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/calculations", handler.Calculate)
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
