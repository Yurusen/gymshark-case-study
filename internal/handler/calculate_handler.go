package handler

import (
	"encoding/json"
	"gymshark-case-study/internal/service"
	"net/http"
)

type CalculateRequest struct {
	Items     int   `json:"items"`
	PackSizes []int `json:"packSizes,omitempty"`
}

type CalculateHandler struct {
	defaultPackSizes []int
}

func NewCalculateHandler(defaultPackSizes []int) *CalculateHandler {
	return &CalculateHandler{
		defaultPackSizes: defaultPackSizes,
	}
}

func (h *CalculateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	origin := r.Header.Get("Origin")

	allowedOrigins := map[string]bool{
		"https://yurusen.github.io":                        true,
		"https://yurusen.github.io/gymshark-case-study-ui": true,
		"http://localhost:3000":                            true,
	}

	if allowedOrigins[origin] {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}

	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request CalculateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil || request.Items < 0 {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	packSizes := h.defaultPackSizes
	if len(request.PackSizes) > 0 {
		packSizes = request.PackSizes
	}

	calculator := service.NewPackCalculator(packSizes)
	result := calculator.CalculatePacks(request.Items)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
