package handler

import (
	"encoding/json"
	"gymshark-case-study/internal/service"
	"net/http"
)

type CalculateRequest struct {
	Items int `json:"items"`
}

type CalculateHandler struct {
	calculator *service.PackCalculator
}

func NewCalculateHandler(calculator *service.PackCalculator) *CalculateHandler {
	return &CalculateHandler{
		calculator: calculator,
	}
}

func (h *CalculateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request CalculateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil || request.Items <= 0 {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	result := h.calculator.CalculatePacks(request.Items)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(result)
}
