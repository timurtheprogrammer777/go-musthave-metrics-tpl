package main

import (
	"fmt"
	"go-musthave-metrics-tpl/internal/runtime"
	"log"
	"strconv"
	"time"

	// "github.com/go-resty/resty/v2"
	"resty.dev/v3"
)

const serverAddress = "http://localhost:8080"

func main() {
	client := resty.New()
	metrics := runtime.NewRuntimeMetrics()

	tickerPoll := time.NewTicker(2 * time.Second)
	tickerReport := time.NewTicker(10 * time.Second)

	for {
		select {
		case <-tickerPoll.C:
			metrics.Collect()
		case <-tickerReport.C:
			sendMetrics(client, metrics)
		}
	}
}

// sendMetrics отправляет метрики на сервер
func sendMetrics(client *resty.Client, metrics *runtime.MetricsRuntime) {
	// Отправляем PollCount
	sendMetric(client, "counter", "PollCount", strconv.FormatInt(metrics.PollCount, 10))

	// Отправляем RandomValue
	sendMetric(client, "gauge", "RandomValue", strconv.FormatFloat(metrics.RandomValue, 'f', 6, 64))
}

// sendMetric отправляет одну метрику
func sendMetric(client *resty.Client, metricType, name, value string) {
	url := fmt.Sprintf("%s/update/%s/%s/%s", serverAddress, metricType, name, value)

	resp, err := client.R().
		SetHeader("Content-Type", "text/plain").
		Post(url)

	if err != nil {
		log.Println("Ошибка отправки метрики:", err)
		return
	}

	log.Printf("Метрика %s (%s) = %s отправлена, статус: %d\n", name, metricType, value, resp.StatusCode())
}
