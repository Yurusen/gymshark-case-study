package main

import (
	"gymshark-case-study/internal/config"
	"gymshark-case-study/internal/handler"
	"gymshark-case-study/internal/service"
	"log"
	"net/http"
)

func main() {
	packSizes := config.LoadPackSizes()
	calculator := service.NewPackCalculator(packSizes)
	calculateHandler := handler.NewCalculateHandler(calculator)

	http.Handle("/calculate", calculateHandler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
