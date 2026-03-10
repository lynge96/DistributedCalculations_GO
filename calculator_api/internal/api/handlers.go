package api

import (
	"calculator_api/internal/models"
	"encoding/json"
	"net/http"
)

// Interface angiver hvad handleren har brug for.
// Interfaces defineres der hvor de bruges, ikke hvor de implementeres
// I dette tilfælde skal den bruge en service der kan Calculate(string)
type CalculatorService interface {
	Calculate(string) (models.CalculationResult, error)
}

type Publisher interface {
	Publish(message models.CalculationResult) error
}

// Handleren gemmer en reference til den service den skal bruge
// Ækvivalent til dependency injection i C#
type Handler struct {
	service   CalculatorService
	publisher Publisher
}

// Laver ny handler og injicerer servicen og returnerer pointer til den
func NewHandler(service CalculatorService, publisher Publisher) *Handler {
	return &Handler{
		service:   service,
		publisher: publisher,
	}
}

type CalculationRequest struct {
	Expression string `json:"expression"`
}

func (h *Handler) Calculate(w http.ResponseWriter, r *http.Request) {

	request, err := decodeCalculationRequest(r)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	result, err := h.service.Calculate(request.Expression)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.publisher.Publish(result); err != nil {
		http.Error(w, "failed to publish calculation result", http.StatusInternalServerError)
		return
	}

	if err := writeJSON(w, result); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func decodeCalculationRequest(r *http.Request) (CalculationRequest, error) {
	var request CalculationRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return CalculationRequest{}, err
	}
	return request, nil
}

func writeJSON(w http.ResponseWriter, value any) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(value)
}
