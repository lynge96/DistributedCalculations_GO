package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"shared/helpers"
	"shared/models"
)

// Interface angiver hvad handleren har brug for.
// Interfaces defineres der hvor de bruges, ikke hvor de implementeres
// I dette tilfælde skal den bruge en service der kan Calculate(string)
type CalculatorService interface {
	Calculate(string) (models.CalculationResult, error)
}

// Handleren gemmer en reference til den service den skal bruge
// Ækvivalent til dependency injection i C#
type Handler struct {
	service CalculatorService
}

// Laver ny handler og injicerer servicen og returnerer pointer til den
func NewHandler(service CalculatorService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Calculate(w http.ResponseWriter, r *http.Request) {

	request, err := decodeCalculationRequest(r)
	if err != nil {
		slog.Warn("failed to decode request", "error", err)
		helpers.Respond(w, http.StatusBadRequest, models.NewCalculationError(request.Expression, err))
		return
	}

	if err := request.Validate(); err != nil {
		slog.Warn("invalid request", "error", err, "expression", request.Expression)
		helpers.Respond(w, http.StatusBadRequest, models.NewCalculationError(request.Expression, err))
		return
	}

	result, err := h.service.Calculate(request.Expression)
	if err != nil {
		slog.Warn("calculation failed", "error", err, "expression", request.Expression)
		helpers.Respond(w, http.StatusBadRequest, models.NewCalculationError(request.Expression, err))
		return
	}

	helpers.Respond(w, http.StatusOK, result)
}

func decodeCalculationRequest(r *http.Request) (models.CalculationRequest, error) {

	var request models.CalculationRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return models.CalculationRequest{}, err
	}
	return request, nil
}
