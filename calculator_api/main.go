package main

import (
	"calculator_api/internal/messaging/rabbitmq"
	"log/slog"
	"net/http"
	"os"
	"shared/auth"
	"shared/configuration"
	"shared/logger"

	"calculator_api/internal/api"
	"calculator_api/internal/calculator"
)

func main() {

	// env vars
	connString := configuration.GetEnv("RABBITMQ_URL", "amqp://guest:guest@raspberrypi:5672/")
	queue := configuration.GetEnv("RABBITMQ_QUEUE", "calculations")
	port := configuration.GetEnv("PORT", "8080")
	secretKey := configuration.GetEnv("JWT_SECRET", "default-secret-key")

	// setup
	logger.Setup()
	publisher, err := rabbitmq.NewPublisher(connString, queue)
	if err != nil {
		slog.Error("failed to create publisher:", "error", err, "queue", queue, "connString", connString)
		os.Exit(1)
	}
	defer publisher.Close()

	parser := &calculator.GovaluateParser{}
	service := calculator.NewService(parser, publisher)
	jwtAuth := auth.NewJwtAuth(secretKey)

	handler := api.NewHandler(service)
	router := api.NewRouter(handler, jwtAuth)

	slog.Info("server running", "port", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		slog.Error("server stopped unexpectedly", "error", err)
		os.Exit(1)
	}
}
