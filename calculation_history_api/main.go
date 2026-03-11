package main

import (
	"history/internal/api"
	"history/internal/consumer/rabbitmq"
	"history/internal/storage"
	"log"
	"net/http"
	"shared/configuration"
)

func main() {

	connString := configuration.GetEnv("RABBITMQ_URL", "amqp://guest:guest@raspberrypi:5672/")
	queue := configuration.GetEnv("RABBITMQ_QUEUE", "calculations")
	port := configuration.GetEnv("PORT", "8081")
	queueSize := configuration.GetEnvInt("RABBITMQ_QUEUE_SIZE", 5)

	historyStore := storage.NewHistoryStore(queueSize)
	handler := api.NewHandler(historyStore)
	router := api.NewRouter(handler)
	consumer, err := rabbitmq.NewConsumer(historyStore, connString, queue)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	defer consumer.Close()

	go func() {
		if err := consumer.Start(); err != nil {
			log.Fatalf("Consumer error: %v", err)
		}
	}()

	log.Printf("Server running on port :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
