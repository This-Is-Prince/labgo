package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Go Weather")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello!")
	})

	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		city := strings.SplitN(r.URL.Path, "/", 3)[2]
		data, err := query(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})
	fmt.Println("Server is listening on port: 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func query(city string) (WeatherData, error) {
	openWeatherMapApiKey := os.Getenv("OPEN_WEATHER_MAP_API_KEY")

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, openWeatherMapApiKey)

	res, err := http.Get(url)
	if err != nil {
		return WeatherData{}, err
	}
	defer res.Body.Close()

	var wd WeatherData
	if err := json.NewDecoder(res.Body).Decode(&wd); err != nil {
		return WeatherData{}, err
	}
	return wd, nil
}
