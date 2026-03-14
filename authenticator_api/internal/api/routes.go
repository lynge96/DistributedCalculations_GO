package api

import (
	"net/http"
	"shared/middleware"
)

func NewRouter(h *Handler, validator middleware.TokenValidator) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/register", h.Register)
	mux.HandleFunc("POST /api/login", h.Login)
	mux.Handle("POST /api/logout", middleware.AuthMiddleware(validator, http.HandlerFunc(h.Logout)))

	return mux
}
