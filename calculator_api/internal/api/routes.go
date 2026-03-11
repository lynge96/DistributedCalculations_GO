package api

import "net/http"

func NewRouter(h *Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/calculations", h.Calculate)

	return mux
}
