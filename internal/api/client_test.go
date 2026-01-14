package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetForecast(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"latitude": 1,
			"longitude": 2,
			"current": {
				"time": "2024-01-01T00:00",
				"temperature_2m": 12.5,
				"wind_speed_10m": 5.4
			}
		}`))
	}))
	defer server.Close()

	resp, err := GetForecast(server.URL, 1, 2)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if resp.Current.Temperature != 12.5 {
		t.Fatalf("expected 12.5, got %v", resp.Current.Temperature)
	}
}
