package main

import (
	"fmt"
	"log"
	"net/http"

	handler "go-musthave-metrics-tpl/internal/handlers"
	"go-musthave-metrics-tpl/internal/storage"
)

func main() {

	storage := storage.NewMemStorage()

	http.HandleFunc("/update/", handler.UpdateHandler(storage))

	fmt.Println("Starting server on port :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
