package runtime

import (
	"fmt"
	"testing"
)

func TestMetricsRuntime_Collect(t *testing.T) {
	rm := NewRuntimeMetrics()

	if rm.PollCount != 0 {
		t.Errorf("Expected PollCount to be 0, got %d", rm.PollCount)
	}
	if rm.RandomValue != 0 {
		t.Errorf("Expected RandomValue to be 0, got %f", rm.RandomValue)
	}

	rm.Collect()

	if rm.PollCount != 1 {
		t.Errorf("Expected PollCount to be 1 after Collect, got %d", rm.PollCount)
	}

	if rm.RandomValue < 0 || rm.RandomValue > 100 {
		t.Errorf("Expected RandomValue to be between 0 and 100, got %f", rm.RandomValue)
	}

	if rm.MemStats.Alloc == 0 {
		t.Error("Expected MemStats.Alloc to be non-zero")
	}

	fmt.Println("runtime metrics are ok!!!")
}
