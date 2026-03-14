package api

import (
	"net/http"
	"shared/middleware"
)

func NewRouter(h *Handler, validator middleware.TokenValidator) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /api/history", middleware.AuthMiddleware(validator, http.HandlerFunc(h.History)))
	mux.Handle("DELETE /api/history/clear", middleware.AuthMiddleware(validator, http.HandlerFunc(h.Clear)))

	return mux
}
