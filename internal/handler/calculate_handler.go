package handler

import (
	"calcOnGO/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type CalculateRequest struct {
	Expression string `json:"expression"`
}

type CalculateResponse struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func Calculate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"Method Not Allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req CalculateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid request format"}`, http.StatusUnprocessableEntity)
		return
	}

	// Вызов функции вычисления
	result, err := service.Calc(req.Expression)
	if err != nil {
		// Проверка типа ошибки
		if err.Error() == "ошибка - деление на ноль" {
			http.Error(w, `{"error":"Division by zero"}`, http.StatusUnprocessableEntity)
			return
		}
		// В случае других ошибок
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusUnprocessableEntity)
		return
	}

	response := CalculateResponse{Result: fmt.Sprintf("%g", result)}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
