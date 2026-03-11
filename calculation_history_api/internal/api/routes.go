package api

import "net/http"

func NewRouter(h *Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/history", h.History)
	mux.HandleFunc("DELETE /api/history/clear", h.Clear)

	return mux
}
