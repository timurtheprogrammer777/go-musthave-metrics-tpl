package main

import (
	"fmt"
	"log"
	"net/http"

	"go-musthave-metrics-tpl/internal/handlers"
	"go-musthave-metrics-tpl/internal/storage"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

const PORT = ":8080"

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	memStorage := storage.NewMemStorage()
	r.Get("/", mainPageMetricNameValue(memStorage))

	r.Get("/value/{type}/{name}", handlers.UpdateHandler(memStorage))

	fmt.Printf("Starting server on port %s", PORT)
	log.Fatal(http.ListenAndServe(PORT, r))
}

func mainPageMetricNameValue(storage *storage.MemStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		// Генерируем HTML с метриками
		html := `<!DOCTYPE html>
		<html lang="ru">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Метрики</title>
			<style>
				body { font-family: Arial, sans-serif; padding: 20px; }
				h1 { color: #333; }
				table { width: 100%%; border-collapse: collapse; margin-top: 20px; }
				th, td { border: 1px solid #ddd; padding: 8px; text-align: left; }
				th { background-color: #f2f2f2; }
			</style>
		</head>
		<body>
			<h1>Метрики сервера</h1>
			<table>
				<tr>
					<th>Тип</th>
					<th>Имя</th>
					<th>Значение</th>
				</tr>`

		// Добавляем данные в HTML
		for key, value := range storage.Gauges {
			html += fmt.Sprintf("<tr><td>gauge</td><td>%s</td><td>%f</td></tr>", key, value)
		}
		for key, value := range storage.Counters {
			html += fmt.Sprintf("<tr><td>counter</td><td>%s</td><td>%d</td></tr>", key, value)
		}

		html += `</table></body></html>`

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, html)
	}
}
