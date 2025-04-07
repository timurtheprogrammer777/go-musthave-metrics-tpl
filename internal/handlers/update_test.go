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
		{"Valid Gauge", "POST", "/update/gauge/testMetric/42.5", http.StatusOK},
		{"Valid Counter", "POST", "/update/counter/testCounter/10", http.StatusOK},
		{"Invalid Method", "GET", "/update/gauge/testMetric/42.5", http.StatusMethodNotAllowed},
		{"Invalid URL Format", "POST", "/update/gauge/testMetric", http.StatusNotFound},
		{"Invalid Gauge Value", "POST", "/update/gauge/testMetric/abc", http.StatusBadRequest},
		{"Invalid Counter Value", "POST", "/update/counter/testCounter/xyz", http.StatusBadRequest},
		{"Invalid Metric Type", "POST", "/update/unknown/testMetric/42", http.StatusBadRequest},
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
