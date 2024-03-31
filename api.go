package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

/*
{
  "latitude": 52.5,
  "longitude": 13.419998,
  "generationtime_ms": 0.17595291137695312,
  "utc_offset_seconds": 0,
  "timezone": "GMT",
  "timezone_abbreviation": "GMT",
  "elevation": 40.0,
  "current_units": {
    "time": "iso8601",
    "interval": "seconds",
    "temperature_2m": "°C",
    "relative_humidity_2m": "%",
    "apparent_temperature": "°C",
    "is_day": "",
    "precipitation": "mm",
    "rain": "mm",
    "showers": "mm",
    "snowfall": "cm",
    "weather_code": "wmo code",
    "cloud_cover": "%",
    "pressure_msl": "hPa",
    "surface_pressure": "hPa",
    "wind_speed_10m": "km/h",
    "wind_direction_10m": "°",
    "wind_gusts_10m": "km/h"
  },
  "current": {
    "time": "2024-03-31T12:45",
    "interval": 900,
    "temperature_2m": 19.2,
    "relative_humidity_2m": 57,
    "apparent_temperature": 19.2,
    "is_day": 1,
    "precipitation": 0.00,
    "rain": 0.00,
    "showers": 0.00,
    "snowfall": 0.00,
    "weather_code": 1,
    "cloud_cover": 9,
    "pressure_msl": 997.6,
    "surface_pressure": 993.0,
    "wind_speed_10m": 2.3,
    "wind_direction_10m": 252,
    "wind_gusts_10m": 15.8
  }
}
*/

type currentWeather struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	GenerationtimeMs     float64 `json:"generationtime_ms"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation            float64 `json:"elevation"`
	// CurrentUnits describes the unit type for each field
	CurrentUnits struct {
		Time                string `json:"time"`
		Interval            string `json:"interval"`
		Temperature2M       string `json:"temperature_2m"`
		RelativeHumidity2M  string `json:"relative_humidity_2m"`
		ApparentTemperature string `json:"apparent_temperature"`
		IsDay               string `json:"is_day"`
		Precipitation       string `json:"precipitation"`
		Rain                string `json:"rain"`
		Showers             string `json:"showers"`
		Snowfall            string `json:"snowfall"`
		WeatherCode         string `json:"weather_code"`
		CloudCover          string `json:"cloud_cover"`
		PressureMsl         string `json:"pressure_msl"`
		SurfacePressure     string `json:"surface_pressure"`
		WindSpeed10M        string `json:"wind_speed_10m"`
		WindDirection10M    string `json:"wind_direction_10m"`
		WindGusts10M        string `json:"wind_gusts_10m"`
	} `json:"current_units"`
	Current struct {
		// Time of obeservation, iso8601
		Time string `json:"time"`
		// Interval for weather data (?), seconds
		Interval            int     `json:"interval"`
		Temperature2M       float64 `json:"temperature_2m"`
		RelativeHumidity2M  int     `json:"relative_humidity_2m"`
		ApparentTemperature float64 `json:"apparent_temperature"`
		IsDay               int     `json:"is_day"`
		Precipitation       float64 `json:"precipitation"`
		Rain                float64 `json:"rain"`
		Showers             float64 `json:"showers"`
		Snowfall            float64 `json:"snowfall"`
		WeatherCode         int     `json:"weather_code"`
		CloudCover          int     `json:"cloud_cover"`
		PressureMsl         float64 `json:"pressure_msl"`
		SurfacePressure     float64 `json:"surface_pressure"`
		WindSpeed10M        float64 `json:"wind_speed_10m"`
		WindDirection10M    int     `json:"wind_direction_10m"`
		WindGusts10M        float64 `json:"wind_gusts_10m"`
	} `json:"current"`
}

func fetchWeather(ctx context.Context, lat, lng float64) (*currentWeather, error) {
	hc := http.DefaultClient

	q := url.Values{}
	q.Add("latitude", fmt.Sprintf("%f", lat))
	q.Add("longitude", fmt.Sprintf("%f", lng))
	// get all the current values. not dealing with forecasts right now
	q.Add("current", "temperature_2m,relative_humidity_2m,apparent_temperature,is_day,precipitation,rain,showers,snowfall,weather_code,cloud_cover,pressure_msl,surface_pressure,wind_speed_10m,wind_direction_10m,wind_gusts_10m")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/v1/forecast?%s", baseURL, q.Encode()), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := hc.Do(req)
	if err != nil {
		return nil, fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected http status code: %d", resp.StatusCode)
	}

	var cw currentWeather
	if err := json.NewDecoder(resp.Body).Decode(&cw); err != nil {
		return nil, fmt.Errorf("reading and decoding response body: %w", err)
	}

	return &cw, nil
}
