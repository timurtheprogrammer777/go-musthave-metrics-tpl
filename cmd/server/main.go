package main

import (
	"fmt"
	"log"
	"net/http"

	"go-musthave-metrics-tpl/internal/handlers"
	"go-musthave-metrics-tpl/internal/storage"
)

func main() {

	memStorage := storage.NewMemStorage()

	http.HandleFunc("/update/", handlers.UpdateHandler(memStorage))

	fmt.Println("Starting server on port :8080")
	fmt.Println("Starting server on port :8080")
	fmt.Println("Starting server on port :8080")
	fmt.Println("Starting server on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
