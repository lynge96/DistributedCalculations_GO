package helpers

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func Respond(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		slog.Warn("failed to encode response:", "error", err, "body", body, "status", status)
	}
}
