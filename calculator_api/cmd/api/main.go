package main

import (
	"log"
	"net/http"

	"calculator_api/internal/api"
	"calculator_api/internal/calculator"
)

func main() {

	parser := &calculator.GovaluateParser{}
	service := calculator.NewService(parser)
	handler := api.NewHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/calculations", handler.Calculate)
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
