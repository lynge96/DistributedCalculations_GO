package api

import (
	"net/http"
	"shared/helpers"
	"shared/models"
)

type HistoryStore interface {
	Clear()
	GetAll() []models.CalculationResult
}

type Handler struct {
	store HistoryStore
}

func NewHandler(store HistoryStore) *Handler {
	return &Handler{store: store}
}

// GET /history
func (h *Handler) History(w http.ResponseWriter, r *http.Request) {

	entries := h.store.GetAll()
	helpers.Respond(w, http.StatusOK, entries)
}

// DELETE /history/clear
func (h *Handler) Clear(w http.ResponseWriter, r *http.Request) {

	h.store.Clear()
	helpers.Respond(w, http.StatusOK, nil)
}
