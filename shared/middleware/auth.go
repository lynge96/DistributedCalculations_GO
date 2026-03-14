package middleware

import (
	"net/http"
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
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(header, "Bearer ")

		_, err := validator.ValidateToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
