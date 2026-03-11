package main

import (
	"history/internal/api"
	"history/internal/consumer/rabbitmq"
	"history/internal/storage"
	"log"
	"net/http"
)

func main() {

	historyStore := storage.NewHistoryStore(5)
	handler := api.NewHandler(historyStore)
	router := api.NewRouter(handler)
	consumer, err := rabbitmq.NewConsumer(
		historyStore,
		"amqp://guest:guest@raspberrypi:5672/",
		"calculations")
	if err != nil {
		log.Fatal(err)
	}

	go consumer.Start()

	log.Println("Server running on :8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
