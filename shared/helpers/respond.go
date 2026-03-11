package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

func Respond(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("failed to encode response: %v", err)
	}
}
