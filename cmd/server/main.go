package main

import (
	"fmt"
	"log"
	"net/http"

	"go-musthave-metrics-tpl/internal/handlers"
	"go-musthave-metrics-tpl/internal/storage"

	"github.com/go-chi/chi/v5"
)

func main() {
	memStorage := storage.NewMemStorage()
	router := chi.NewRouter()

	router.Post("/update/{metricType}/{name}/{value}", handlers.UpdateHandler(memStorage))
	router.Get("/", handlers.UpdateMainPage(memStorage))

	fmt.Println("Starting server on port :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
