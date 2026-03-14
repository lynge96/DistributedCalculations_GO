package main

import (
	"authenticator_api/internal/api"
	"authenticator_api/internal/auth"
	"authenticator_api/internal/service"
	"authenticator_api/internal/storage"
	"log/slog"
	"net/http"
	"os"
	sharedAuth "shared/auth"
	"shared/configuration"
	"shared/logger"

	_ "modernc.org/sqlite"
)

func main() {

	// env vars
	secretKey := configuration.GetEnv("JWT_SECRET", "default-secret-key")
	dbPath := configuration.GetEnv("DB_PATH", "./data/users.db")
	port := configuration.GetEnv("PORT", "8082")

	// setup
	logger.Setup()
	jwtAuth := sharedAuth.NewJwtAuth(secretKey)
	authenticator := auth.NewAuthenticator(jwtAuth)
	userStore, err := storage.NewUserStore(dbPath)
	if err != nil {
		slog.Error("failed to create user store", "error", err)
		os.Exit(1)
	}

	authService := service.NewService(userStore, authenticator)
	handler := api.NewHandler(authService)
	router := api.NewRouter(handler)

	slog.Info("server running", "port", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		slog.Error("server stopped unexpectedly", "error", err)
		os.Exit(1)
	}
}
