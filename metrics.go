package main

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const metricsNamespace = "openmeteo"

// their version, hope it's consistent
const iso8601Format = "2006-01-02T15:04"

var (
	obvTime = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: metricsNamespace,
			Name:      "time",
			Help:      "Time of observation, unix format",
		},
		[]string{"location"},
	)

	// these are initialized via the function, to set units appropriately

	interval            *prometheus.GaugeVec
	temperature         *prometheus.GaugeVec
	relativeHumidity    *prometheus.GaugeVec
	apparentTemperature *prometheus.GaugeVec
	isDaytime           *prometheus.GaugeVec
	precipitation       *prometheus.GaugeVec
	rain                *prometheus.GaugeVec
	showers             *prometheus.GaugeVec
	snowfall            *prometheus.GaugeVec
	weatherCode         *prometheus.GaugeVec
	cloudCover          *prometheus.GaugeVec
	pressureMSL         *prometheus.GaugeVec
	surfacePressure     *prometheus.GaugeVec
	windSpeed           *prometheus.GaugeVec
	windDirection       *prometheus.GaugeVec
	windGusts           *prometheus.GaugeVec
)

// initializeMetrics must be called before any metrics are created, to set the
// value units in the help text
func initializeMetrics(cw currentWeather) {
	interval = newGauge("interval", "interval of report", cw.CurrentUnits.Interval)
	temperature = newGauge("temperature", "current temperature (2m)", cw.CurrentUnits.Temperature2M)
	relativeHumidity = newGauge("relative_humidity", "relative humidity (2m)", cw.CurrentUnits.RelativeHumidity2M)
	apparentTemperature = newGauge("apparent_temperature", "apparent temperature", cw.CurrentUnits.ApparentTemperature)
	isDaytime = newGauge("is_day", "Is daytime", "1 when daytime, 0 when not")
	precipitation = newGauge("precipitation", "precipitation", cw.CurrentUnits.Precipitation)
	rain = newGauge("rain", "rain", cw.CurrentUnits.Rain)
	showers = newGauge("showers", "showers", cw.CurrentUnits.Showers)
	snowfall = newGauge("snowfall", "snowfall", cw.CurrentUnits.Snowfall)
	weatherCode = newGauge("weather_code", "???", cw.CurrentUnits.WeatherCode)
	cloudCover = newGauge("cloud_cover", "cloud cover", cw.CurrentUnits.CloudCover)
	pressureMSL = newGauge("pressure_msl", "pressure at MSL", cw.CurrentUnits.PressureMsl)
	surfacePressure = newGauge("pressure_surface", "pressure at surface", cw.CurrentUnits.SurfacePressure)
	windSpeed = newGauge("wind_speed", "wind speed (10m)", cw.CurrentUnits.WindSpeed10M)
	windDirection = newGauge("wind_direction", "direction of wind (10m)", cw.CurrentUnits.WindDirection10M)
	windGusts = newGauge("wind_gusts", "wind gusts (10m)", cw.CurrentUnits.WindGusts10M)
}

func setMetrics(locationName string, cw currentWeather) error {
	if cw.CurrentUnits.Time != "iso8601" {
		return fmt.Errorf("unhandled time format: %s", cw.CurrentUnits.Time)
	}
	ct, err := time.Parse(iso8601Format, cw.Current.Time)
	if err != nil {
		return fmt.Errorf("parsing %s time %s with %s: %w", cw.CurrentUnits.Time, cw.Current.Time, iso8601Format, err)
	}

	lbl := prometheus.Labels{"location": locationName}

	obvTime.With(lbl).Set(float64(ct.Unix()))
	interval.With(lbl).Set(float64(cw.Current.Interval))
	temperature.With(lbl).Set(cw.Current.Temperature2M)
	relativeHumidity.With(lbl).Set(float64(cw.Current.RelativeHumidity2M))
	apparentTemperature.With(lbl).Set(cw.Current.ApparentTemperature)
	isDaytime.With(lbl).Set(float64(cw.Current.IsDay))
	precipitation.With(lbl).Set(cw.Current.Precipitation)
	rain.With(lbl).Set(cw.Current.Rain)
	showers.With(lbl).Set(cw.Current.Showers)
	snowfall.With(lbl).Set(cw.Current.Snowfall)
	weatherCode.With(lbl).Set(float64(cw.Current.WeatherCode))
	cloudCover.With(lbl).Set(float64(cw.Current.CloudCover))
	pressureMSL.With(lbl).Set(cw.Current.PressureMsl)
	surfacePressure.With(lbl).Set(cw.Current.SurfacePressure)
	windSpeed.With(lbl).Set(cw.Current.WindSpeed10M)
	windDirection.With(lbl).Set(float64(cw.Current.WindDirection10M))
	windGusts.With(lbl).Set(cw.Current.WindGusts10M)

	return nil
}

func newGauge(name, help, format string) *prometheus.GaugeVec {
	return promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: metricsNamespace,
			Name:      name,
			Help:      fmt.Sprintf("%s, %s", help, format),
		},
		[]string{"location"},
	)
}
