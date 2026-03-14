package middleware

import (
	"log/slog"
	"net/http"
	"shared/helpers"
	"shared/models"
	"strings"
)

type TokenValidator interface {
	ValidateToken(tokenString string) (*models.Claims, error)
}

func AuthMiddleware(validator TokenValidator, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		header := r.Header.Get("Authorization")
		if header == "" {
			slog.Warn("missing authorization header")
			helpers.Respond(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		token := strings.TrimPrefix(header, "Bearer ")

		_, err := validator.ValidateToken(token)
		if err != nil {
			slog.Warn("invalid token", "error", err)
			helpers.Respond(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		next.ServeHTTP(w, r)
	})
}
