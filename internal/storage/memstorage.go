package storage

import (
	"fmt"
)

type MemStorage struct {
	Gauges   map[string]float64
	Counters map[string]int64
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		Gauges:   make(map[string]float64),
		Counters: make(map[string]int64),
	}
}

func (s *MemStorage) UpdateGauge(name string, value float64) {
	s.Gauges[name] = value
	fmt.Println(s.Gauges)
}

func (s *MemStorage) UpdateCounter(name string, value int64) {
	s.Counters[name] += value
	fmt.Println(s.Counters)
}

func (s *MemStorage) GetAllGauges() map[string]float64 {
	return s.Gauges
}

func (s *MemStorage) GetAllCounters() map[string]int64 {
	return s.Counters
}

func (s *MemStorage) GetAllMetrics() map[string]interface{} {
	metrics := make(map[string]interface{})

	for name, value := range s.Gauges {
		metrics[name] = value
	}

	for name, value := range s.Counters {
		metrics[name] = value
	}
	return metrics
}
