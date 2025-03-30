package handlers_test

import (
	"bytes"
	"fmt"
	"go-musthave-metrics-tpl/internal/handlers"
	"go-musthave-metrics-tpl/internal/storage"
	"net/http"
	"net/http/httptest"
	"testing"
)

type UpdateTest []struct {
	name         string
	method       string
	url          string
	expectedCode int
}

func TestUpdateHandler(t *testing.T) {
	memStorage := storage.NewMemStorage()
	handler := handlers.UpdateHandler(memStorage)

	tests := UpdateTest{
		{"Valid Gauge", "POST", "/value/gauge/testMetric/42.5", http.StatusOK},
		{"Valid Counter", "POST", "/value/counter/testCounter/10", http.StatusOK},
		{"Invalid Method", "GET", "/value/gauge/testMetric/42.5", http.StatusMethodNotAllowed},
		{"Invalid URL Format", "POST", "/value/gauge/testMetric", http.StatusNotFound},
		{"Invalid Gauge Value", "POST", "/value/gauge/testMetric/abc", http.StatusBadRequest},
		{"Invalid Counter Value", "POST", "/value/counter/testCounter/xyz", http.StatusBadRequest},
		{"Invalid Metric Type", "POST", "/value/unknown/testMetric/42", http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.url, bytes.NewBuffer(nil))
			rec := httptest.NewRecorder()

			handler.ServeHTTP(rec, req)

			if rec.Code != tt.expectedCode {
				t.Errorf("Expected status %d, got %d", tt.expectedCode, rec.Code)
			}
		})
	}
	fmt.Println("Handlers updating tests are ok!")
}
