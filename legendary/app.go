package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const WEATHER_API_KEY = "YOUR_WEATHER_API_KEY"

type Location struct {
	City      string  `json:"city"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type WeatherResponse struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func getLocation(ip string) (string, float64, float64, error) {
	url := fmt.Sprintf("https://ipapi.co/%s/json/", ip)
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, 0, fmt.Errorf("error fetching location: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", 0, 0, fmt.Errorf("received non-200 response: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0, 0, fmt.Errorf("error reading response body: %v", err)
	}

	var location Location
	if err := json.Unmarshal(body, &location); err != nil {
		return "", 0, 0, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	return location.City, location.Latitude, location.Longitude, nil
}

func getWeather(lat, lon float64) (string, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%.2f&lon=%.2f&appid=%s&units=metric", lat, lon, WEATHER_API_KEY)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error fetching weather: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-200 response: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	var weatherData WeatherResponse
	if err := json.Unmarshal(body, &weatherData); err != nil {
		return "", fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	weatherDesc := weatherData.Weather[0].Description
	temperature := weatherData.Main.Temp
	return fmt.Sprintf("Weather: %s, Temperature: %.2f°C", weatherDesc, temperature), nil
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		clientIP := c.ClientIP()
		city, lat, lon, err := getLocation(clientIP)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to determine location from IP address"})
			return
		}

		weatherInfo, err := getWeather(lat, lon)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch weather data"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"ip":      clientIP,
			"city":    city,
			"weather": weatherInfo,
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
