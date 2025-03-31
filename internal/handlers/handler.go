package handler

import (
	"fmt"
	"go-musthave-metrics-tpl/internal/storage"
	"net/http"
	"strconv"
	"strings"
)

func UpdateHandler(storage *storage.MemStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/update/"), "/")

		if len(parts) != 3 {
			http.Error(w, "Invalid URL format", http.StatusNotFound)
			return
		}

		metricType, metricnName, metricValue := parts[0], parts[1], parts[2]

		switch metricType {
		case "gauge":
			value, err := strconv.ParseFloat(metricValue, 64)
			if err != nil {
				http.Error(w, "Invalid gauge value", http.StatusBadRequest)
				return
			}
			storage.UpdateGauge(metricnName, value)
		case "counter":
			value, err := strconv.ParseFloat(metricValue, 64)
			if err != nil {
				http.Error(w, "Invalid counter value", http.StatusBadRequest)
				return
			}
			storage.UpdateCounter(metricnName, int64(value))
		default:
			http.Error(w, "Invalid metric type", http.StatusBadRequest)
			return
		}
		fmt.Printf("Received metric: Type=%s, Name=%s, Value=%s\n", metricType, metricnName, metricValue)
		w.WriteHeader(http.StatusOK)
	}
}
