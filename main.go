package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type WeatherValues struct {
	ValueWater int32 `json:"value_water"`
	ValueWind  int32 `json:"value_wind"`
}

type WeatherData struct {
	Data        WeatherValues `json:"data"`
	StatusWater string        `json:"status_water"`
	StatusWind  string        `json:"status_wind"`
}

func main() {
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	timer := time.NewTimer(20 * time.Second)
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:

			valueWater := randGen.Int31n(100) + 1
			valueWind := randGen.Int31n(100) + 1
			statusWater := getStatusWater(valueWater)
			statusWind := getStatusWind(valueWind)
			weatherData := WeatherData{
				Data: WeatherValues{
					ValueWater: valueWater,
					ValueWind:  valueWind,
				},
				StatusWater: statusWater,
				StatusWind:  statusWind,
			}
			payload, _ := json.Marshal(weatherData)
			resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(payload))
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Printf("{\n water: %v,\n wind: %v,\n}\nstatus water: %v\nstatus wind: %v\n", weatherData.Data.ValueWater, weatherData.Data.ValueWind, weatherData.StatusWater, weatherData.StatusWind)
			fmt.Printf("Post success: %s\n", resp.Status)

		case <-timer.C:
			fmt.Println("Waktu Loop selesai.")
			return
		}
	}
}

func getStatusWater(valueWater int32) string {
	if valueWater < 5 {
		return "Aman"
	} else if valueWater >= 5 && valueWater <= 8 {
		return "Siaga"
	} else {
		return "Bahaya"
	}
}

func getStatusWind(valueWind int32) string {
	if valueWind < 6 {
		return "Aman"
	} else if valueWind >= 6 && valueWind <= 15 {
		return "Siaga"
	} else {
		return "Bahaya"
	}
}
