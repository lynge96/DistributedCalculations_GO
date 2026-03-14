package api

import (
	"net/http"
	"shared/middleware"
)

func NewRouter(h *Handler, validator middleware.TokenValidator) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("POST /api/calculations", middleware.AuthMiddleware(validator, http.HandlerFunc(h.Calculate)))

	return mux
}
