package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type ForecastResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Current   struct {
		Time        string  `json:"time"`
		Temperature float64 `json:"temperature_2m"`
		WindSpeed   float64 `json:"wind_speed_10m"`
	} `json:"current"`
}

func GetForecast(baseURL string, lat, lon float64) (ForecastResponse, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		return ForecastResponse{}, err
	}

	q := u.Query()
	q.Set("latitude", fmt.Sprintf("%f", lat))
	q.Set("longitude", fmt.Sprintf("%f", lon))
	q.Set("current", "temperature_2m,wind_speed_10m")
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return ForecastResponse{}, err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return ForecastResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return ForecastResponse{}, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var fr ForecastResponse
	if err := json.NewDecoder(resp.Body).Decode(&fr); err != nil {
		return ForecastResponse{}, err
	}

	return fr, nil
}
