package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"go-musthave-metrics-tpl/internal/storage"

	"github.com/go-chi/chi/v5"
)

func UpdateHandler(storage *storage.MemStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		metricType := chi.URLParam(r, "metricType")
		name := chi.URLParam(r, "name")
		valueStr := chi.URLParam(r, "value")

		switch metricType {
		case "gauge":
			value, err := strconv.ParseFloat(valueStr, 64)
			if err != nil {
				http.Error(w, "Invalid gauge value", http.StatusBadRequest)
				return
			}
			storage.UpdateGauge(name, value)

		case "counter":
			value, err := strconv.ParseInt(valueStr, 10, 64)
			if err != nil {
				http.Error(w, "Invalid counter value", http.StatusBadRequest)
				return
			}
			storage.UpdateCounter(name, value)

		default:
			http.Error(w, "Invalid metric type", http.StatusBadRequest)
			return
		}

		log.Printf("Received metric: Type=%s, Name=%s, Value=%s\n", metricType, name, valueStr)
		w.WriteHeader(http.StatusOK)
	}
}

func UpdateMainPage(storage *storage.MemStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		metrics := storage.GetAllMetrics()
		w.Header().Set("Content-Type", "text/html")

		fmt.Fprintf(w, "<html><body><h1>Metrics</h1><ul>")
		for name, value := range metrics {
			fmt.Fprintf(w, "<li>%s: %v</li>", name, value)
		}
		fmt.Fprintf(w, "</ul></body></html>")
	}
}
