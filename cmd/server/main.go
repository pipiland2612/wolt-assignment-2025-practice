package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"golang-api-practice/internal/client"
	"golang-api-practice/internal/handler"
	"golang-api-practice/internal/service"
)

func main() {
	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}

	cli := client.NewClient(httpClient)
	svc := service.NewService(cli)
	deliveryHandler := handler.NewHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/delivery-order-price", deliveryHandler.DeliveryPrice)

	server := &http.Server{
		Addr:         ":8000",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("Server running on http://localhost:8000")

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("server failed: %v", err)
	}
}
