package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"go-musthave-metrics-tpl/internal/storage"
)

func UpdateHandler(storage *storage.MemStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		path := strings.TrimPrefix(r.URL.Path, "/update/")
		parts := strings.Split(path, "/")

		if len(parts) != 3 {
			http.Error(w, "Invalid URL format", http.StatusNotFound)
			return
		}

		metricType, name, valueStr := parts[0], parts[1], parts[2]

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
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
	}
}
