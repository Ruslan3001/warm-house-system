package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// TemperatureResponse соответствует структуре в smart_home/services/temperature_service.go
type TemperatureResponse struct {
	Value       float64   `json:"value"`       // Исправлено: было temperature
	Unit        string    `json:"unit"`
	Timestamp   time.Time `json:"timestamp"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	SensorID    string    `json:"sensor_id"`   // Исправлено: было id
	SensorType  string    `json:"sensor_type"`
	Description string    `json:"description"`
}

func main() {
	// Инициализация генератора случайных чисел
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/temperature", func(w http.ResponseWriter, r *http.Request) {
		// Читаем параметры запроса
		location := r.URL.Query().Get("location")
		sensorID := r.URL.Query().Get("sensor_id")

		// Логика из задания: Если нет location, определяем по ID
		if location == "" {
			switch sensorID {
			case "1":
				location = "Living Room"
			case "2":
				location = "Bedroom"
			case "3":
				location = "Kitchen"
			default:
				location = "Unknown"
			}
		}

		// Логика из задания: Если нет ID, определяем по location
		if sensorID == "" {
			switch location {
			case "Living Room":
				sensorID = "1"
			case "Bedroom":
				sensorID = "2"
			case "Kitchen":
				sensorID = "3"
			default:
				sensorID = "0"
			}
		}

		// Генерация случайной температуры (18.0 - 28.0)
		randomTemp := 18.0 + rand.Float64()*(28.0-18.0)

		response := TemperatureResponse{
			Value:       randomTemp,
			Unit:        "Celsius",
			Timestamp:   time.Now(),
			Location:    location,
			Status:      "active",
			SensorID:    sensorID,
			SensorType:  "temperature",
			Description: fmt.Sprintf("Sensor in %s", location),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	fmt.Println("Temperature API is running on port 8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
