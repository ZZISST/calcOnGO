package main

import (
	"calcOnGO/internal/handler"
	"log"
	"net/http"
)

func main() {
	// Регистрация маршрута
	http.HandleFunc("/api/v1/calculate", handler.Calculate)

	// Запуск сервера
	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
