package runtime_test

import (
	"go-musthave-metrics-tpl/internal/runtime" // Убедись, что импорт правильный
	"testing"
)

func TestMetricsRuntime_Collect(t *testing.T) {
	rm := runtime.NewRuntimeMetrics()

	// Проверка начальных значений
	if rm.PollCount != 0 {
		t.Errorf("Expected PollCount to be 0, got %d", rm.PollCount)
	}
	if rm.RandomValue != 0 {
		t.Errorf("Expected RandomValue to be 0, got %f", rm.RandomValue)
	}

	// Выполняем сбор метрик
	rm.Collect()

	// Проверка, что PollCount увеличился
	if rm.PollCount != 1 {
		t.Errorf("Expected PollCount to be 1 after Collect, got %d", rm.PollCount)
	}

	// Проверка, что RandomValue в допустимом диапазоне
	if rm.RandomValue < 0 || rm.RandomValue > 100 {
		t.Errorf("Expected RandomValue to be between 0 and 100, got %f", rm.RandomValue)
	}

	// Проверка использования памяти: MemStats.Alloc должно быть больше 0
	if rm.MemStats.Alloc == 0 {
		t.Error("Expected MemStats.Alloc to be non-zero after Collect")
	}

	// Проверка других ключевых значений MemStats (например, TotalAlloc)
	if rm.MemStats.TotalAlloc == 0 {
		t.Error("Expected MemStats.TotalAlloc to be non-zero after Collect")
	}

	// Также можем проверить, что количество доступной памяти увеличилось или уменьшилось
	// в зависимости от изменений, происходящих в системе.

	t.Log("runtime metrics are ok!")
}
