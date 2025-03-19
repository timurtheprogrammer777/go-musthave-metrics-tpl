package handlers_test

import (
	"bytes"
	"go-musthave-metrics-tpl/internal/handlers"
	"go-musthave-metrics-tpl/internal/storage"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateHandler(t *testing.T) {
	memStorage := storage.NewMemStorage()
	handler := handlers.UpdateHandler(memStorage)

	tests := []struct {
		name         string
		method       string
		url          string
		expectedCode int
	}{
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
}

func TestUpdateMainPage(t *testing.T) {
	memStorage := storage.NewMemStorage()

	// Добавляем метрики в хранилище
	memStorage.UpdateGauge("testMetric", 42.5)
	memStorage.UpdateCounter("testCounter", 10)

	// Тестируем страницу с метриками
	t.Run("Get all metrics", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		rec := httptest.NewRecorder()

		handler := handlers.UpdateMainPage(memStorage)
		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, rec.Code)
		}

		expectedResponse := "<html><body><h1>Metrics</h1><ul><li>testMetric: 42.5</li><li>testCounter: 10</li></ul></body></html>"
		if rec.Body.String() != expectedResponse {
			t.Errorf("Expected response body to be '%s', got '%s'", expectedResponse, rec.Body.String())
		}
	})
}
