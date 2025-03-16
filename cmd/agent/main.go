package main

import (
	"bytes"
	"fmt"
	"go-musthave-metrics-tpl/internal/runtime"
	"log"
	"net/http"
	"strconv"
	"time"
)

const serverAddress = "http://localhost:8080"

func main() {
	metrics := runtime.NewRuntimeMetrics()

	tickerPoll := time.NewTicker(2 * time.Second)
	tickerReport := time.NewTicker(10 * time.Second)

	for {
		select {
		case <-tickerPoll.C:
			metrics.Collect()
		case <-tickerReport.C:
			sendMetrics(metrics)
		}
	}
}

// sendMetrics отправляет метрики на сервер
func sendMetrics(metrics *runtime.MetricsRuntime) {
	client := &http.Client{}

	// Отправляем PollCount
	sendMetric(client, "counter", "PollCount", strconv.FormatInt(metrics.PollCount, 10))

	// Отправляем RandomValue
	sendMetric(client, "gauge", "RandomValue", strconv.FormatFloat(metrics.RandomValue, 'f', 6, 64))
}

// sendMetric отправляет одну метрику
func sendMetric(client *http.Client, metricType, name, value string) {
	url := fmt.Sprintf("%s/update/%s/%s/%s", serverAddress, metricType, name, value)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(nil))
	if err != nil {
		log.Println("Ошибка создания запроса:", err)
		return
	}

	req.Header.Set("Content-Type", "text/plain")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Ошибка отправки метрики:", err)
		return
	}
	defer resp.Body.Close()

	log.Printf("Метрика %s (%s) = %s отправлена, статус: %d\n", name, metricType, value, resp.StatusCode)
}
