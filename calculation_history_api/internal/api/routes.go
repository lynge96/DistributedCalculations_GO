package api

import "net/http"

func NewRouter(h *Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /history", h.History)
	mux.HandleFunc("POST /history/clear", h.Clear)

	return mux
}
