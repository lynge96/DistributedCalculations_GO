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
		respond(w, http.StatusBadRequest, models.CalculationResult{Error: "invalid request"})
		return
	}

	result, err := h.service.Calculate(request.Expression)
	if err != nil {
		respond(w, http.StatusBadRequest, models.CalculationResult{Error: err.Error()})
		return
	}

	if err := h.publisher.Publish(result); err != nil {
		respond(w, http.StatusInternalServerError, models.CalculationResult{Error: err.Error()})
		return
	}

	respond(w, http.StatusOK, result)
}

func decodeCalculationRequest(r *http.Request) (CalculationRequest, error) {

	var request CalculationRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return CalculationRequest{}, err
	}
	return request, nil
}

func respond(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		return
	}
}
