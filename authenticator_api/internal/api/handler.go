package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"shared/helpers"
	"time"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type AuthService interface {
	CreateUser(username, password string) error
	Login(username, password string) (string, error)
}

type Handler struct {
	auth AuthService
}

func NewHandler(auth AuthService) *Handler {
	return &Handler{auth: auth}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {

	var request AuthRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		slog.Error("failed to decode request", "error", err)
		helpers.Respond(w, http.StatusBadRequest, "invalid request")
		return
	}

	err = h.auth.CreateUser(request.Username, request.Password)
	if err != nil {
		slog.Error("failed to create user", "error", err)
		helpers.Respond(w, http.StatusInternalServerError, "failed to create user")
		return
	}

	slog.Info("user registered", "username", request.Username)
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	var request AuthRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		slog.Error("failed to decode request", "error", err)
		helpers.Respond(w, http.StatusBadRequest, "invalid request")
		return
	}

	token, err := h.auth.Login(request.Username, request.Password)
	if err != nil {
		slog.Error("failed to login", "error", err)
		helpers.Respond(w, http.StatusInternalServerError, "failed to login")
		return
	}

	slog.Info("user logged in", "username", request.Username)
	helpers.Respond(w, http.StatusOK, LoginResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   int((24 * time.Hour).Seconds()),
	})
}
