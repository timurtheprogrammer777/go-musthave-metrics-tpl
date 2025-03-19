package storage_test

import (
	"go-musthave-metrics-tpl/internal/storage"
	"testing"
)

func TestMemStorage_UpdateGauge(t *testing.T) {
	s := storage.NewMemStorage()

	// Проверяем, что изначально карты пусты
	if len(s.Gauges) != 0 {
		t.Errorf("Expected Gauges map to be empty, got %v", s.Gauges)
	}

	// Обновляем значение гауже
	s.UpdateGauge("temperature", 23.5)

	// Проверяем, что значение записано правильно
	if s.Gauges["temperature"] != 23.5 {
		t.Errorf("Expected temperature to be 23.5, got %f", s.Gauges["temperature"])
	}

	// Проверяем, что другие метрики не были изменены
	if len(s.Gauges) != 1 {
		t.Errorf("Expected 1 gauge, got %d", len(s.Gauges))
	}

	t.Log("storage gauge updated ok!")
}

func TestMemStorage_UpdateCounter(t *testing.T) {
	s := storage.NewMemStorage()

	// Проверяем, что изначально карты пусты
	if len(s.Counters) != 0 {
		t.Errorf("Expected Counters map to be empty, got %v", s.Counters)
	}

	// Обновляем значение счетчика
	s.UpdateCounter("requests", 5)

	// Проверяем, что значение записано правильно
	if s.Counters["requests"] != 5 {
		t.Errorf("Expected requests to be 5, got %d", s.Counters["requests"])
	}

	// Обновляем снова, чтобы проверить добавление
	s.UpdateCounter("requests", 3)

	// Проверяем, что значение счетчика корректно обновляется
	if s.Counters["requests"] != 8 {
		t.Errorf("Expected requests to be 8 after update, got %d", s.Counters["requests"])
	}

	// Проверяем, что другие метрики не были изменены
	if len(s.Counters) != 1 {
		t.Errorf("Expected 1 counter, got %d", len(s.Counters))
	}

	t.Log("storage counter updated ok!")
}

func TestMemStorage_GetAllMetrics(t *testing.T) {
	s := storage.NewMemStorage()

	// Обновляем метрики
	s.UpdateGauge("temperature", 23.5)
	s.UpdateCounter("requests", 5)

	// Получаем все метрики
	metrics := s.GetAllMetrics()

	// Проверяем, что метрики правильно собраны
	if len(metrics) != 2 {
		t.Errorf("Expected 2 metrics, got %d", len(metrics))
	}

	// Проверяем, что каждая метрика имеет корректные значения
	if value, ok := metrics["temperature"]; !ok || value != 23.5 {
		t.Errorf("Expected 'temperature' to be 23.5, got %v", value)
	}
	if value, ok := metrics["requests"]; !ok || value != 8 {
		t.Errorf("Expected 'requests' to be 8, got %v", value)
	}

	t.Log("storage get all metrics ok!")
}

func TestMemStorage_GetAllGauges(t *testing.T) {
	s := storage.NewMemStorage()

	// Обновляем метрики
	s.UpdateGauge("temperature", 23.5)
	s.UpdateGauge("humidity", 60.0)

	// Получаем все гауже
	gauges := s.GetAllGauges()

	// Проверяем, что полученные гауже корректны
	if len(gauges) != 2 {
		t.Errorf("Expected 2 gauges, got %d", len(gauges))
	}

	// Проверяем, что значения гауже корректны
	if value, ok := gauges["temperature"]; !ok || value != 23.5 {
		t.Errorf("Expected 'temperature' gauge to be 23.5, got %v", value)
	}
	if value, ok := gauges["humidity"]; !ok || value != 60.0 {
		t.Errorf("Expected 'humidity' gauge to be 60.0, got %v", value)
	}

	t.Log("storage get all gauges ok!")
}

func TestMemStorage_GetAllCounters(t *testing.T) {
	s := storage.NewMemStorage()

	// Обновляем метрики
	s.UpdateCounter("requests", 5)
	s.UpdateCounter("errors", 3)

	// Получаем все счетчики
	counters := s.GetAllCounters()

	// Проверяем, что полученные счетчики корректны
	if len(counters) != 2 {
		t.Errorf("Expected 2 counters, got %d", len(counters))
	}

	// Проверяем, что значения счетчиков корректны
	if value, ok := counters["requests"]; !ok || value != 5 {
		t.Errorf("Expected 'requests' counter to be 5, got %v", value)
	}
	if value, ok := counters["errors"]; !ok || value != 3 {
		t.Errorf("Expected 'errors' counter to be 3, got %v", value)
	}

	t.Log("storage get all counters ok!")
}
