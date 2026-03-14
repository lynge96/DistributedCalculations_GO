package main

import (
	"history/internal/api"
	"history/internal/consumer/rabbitmq"
	"history/internal/storage"
	"log/slog"
	"net/http"
	"os"
	"shared/auth"
	"shared/configuration"
	"shared/logger"
)

func main() {

	// env vars
	connString := configuration.GetEnv("RABBITMQ_URL", "amqp://guest:guest@raspberrypi:5672/")
	queue := configuration.GetEnv("RABBITMQ_QUEUE", "calculations")
	port := configuration.GetEnv("PORT", "8081")
	queueSize := configuration.GetEnvInt("RABBITMQ_QUEUE_SIZE", 5)
	secretKey := configuration.GetEnv("JWT_SECRET", "default-secret-key")

	// setup
	logger.Setup()
	historyStore := storage.NewHistoryStore(queueSize)
	handler := api.NewHandler(historyStore)
	jwtAuth := auth.NewJwtAuth(secretKey)
	router := api.NewRouter(handler, jwtAuth)

	consumer, err := rabbitmq.NewConsumer(historyStore, connString, queue)
	if err != nil {
		slog.Error("failed to create consumer", "error", err, "queue", queue, "connString", connString)
		os.Exit(1)
	}
	defer consumer.Close()

	go func() {
		if err := consumer.Start(); err != nil {
			slog.Error("consumer error", "error", err)
			os.Exit(1)
		}
	}()

	slog.Info("server running", "port", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		slog.Error("server stopped unexpectedly", "error", err)
		os.Exit(1)
	}
}
