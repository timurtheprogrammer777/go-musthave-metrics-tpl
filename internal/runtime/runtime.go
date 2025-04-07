package runtime

import (
	"math/rand"
	"runtime"
)

type MetricsRuntime struct {
	PollCount   int64
	RandomValue float64
	MemStats    runtime.MemStats
}

func NewRuntimeMetrics() *MetricsRuntime {
	return &MetricsRuntime{}
}

func (rm *MetricsRuntime) Collect() {
	rm.PollCount++
	rm.RandomValue = rand.Float64() * 100
	runtime.ReadMemStats(&rm.MemStats)
}
