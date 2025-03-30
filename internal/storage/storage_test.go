package storage

import (
	"fmt"
	"testing"
)

func TestMemStorage_UpdateGauge(t *testing.T) {
	s := NewMemStorage()

	if len(s.Gauges) != 0 {
		t.Errorf("Expected Gauges map to be empty, got %v", s.Gauges)
	}

	s.UpdateGauge("temperature", 23.5)

	if s.Gauges["temperature"] != 23.5 {
		t.Errorf("Expected temperature to be 23.5, got %f", s.Gauges["temperature"])
	}

	fmt.Println("storage gauge updated ok!")
}

func TestMemStorage_UpdateCounter(t *testing.T) {
	s := NewMemStorage()

	if len(s.Counters) != 0 {
		t.Errorf("Expected Counters map to be empty, got %v", s.Counters)
	}

	s.UpdateCounter("requests", 5)

	if s.Counters["requests"] != 5 {
		t.Errorf("Expected requests to be 5, got %d", s.Counters["requests"])
	}

	s.UpdateCounter("requests", 3)

	if s.Counters["requests"] != 8 {
		t.Errorf("Expected requests to be 8 after update, got %d", s.Counters["requests"])
	}

	fmt.Println("storage counter updated ok!")

}
