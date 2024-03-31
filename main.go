package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const baseURL = "https://api.open-meteo.com"

func main() {
	var opts struct {
		lat          float64
		lng          float64
		addr         string
		locationName string
		interval     time.Duration
	}
	flag.Float64Var(&opts.lat, "lat", 0, "latitude")
	flag.Float64Var(&opts.lng, "lng", 0, "longitude")
	flag.StringVar(&opts.addr, "addr", "127.0.0.1:8889", "address to listen on")
	flag.StringVar(&opts.locationName, "location-name", "", "friendly location name for metrics, defaults to lat,lng")
	flag.DurationVar(&opts.interval, "interval", 5*time.Minute, "interval to fetch weather")
	flag.Parse()

	var flagErr error
	if opts.lat == 0 {
		flagErr = errors.Join(flagErr, errors.New("lat must be specified, or you must not be on null island"))
	}
	if opts.lat == 0 {
		flagErr = errors.Join(flagErr, errors.New("lng must be specified, or you must not be on null island"))
	}
	if flagErr != nil {
		fmt.Fprint(os.Stderr, flagErr.Error())
		os.Exit(2)
	}

	if opts.locationName == "" {
		opts.locationName = fmt.Sprintf("%f,%f", opts.lat, opts.lng)
	}

	logger := slog.With("location", opts.locationName)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	initw, err := fetchWeather(ctx, opts.lat, opts.lng)
	if err != nil {
		logger.Error("initial weather fetch failed, exiting", "err", err)
		os.Exit(1)
	}

	initializeMetrics(*initw)
	if err := setMetrics(opts.locationName, *initw); err != nil {
		logger.Error("initial metrics set failed, exiting", "err", err)
		os.Exit(1)
	}

	go func() {
		for range time.NewTicker(opts.interval).C {
			logger.Info("fetching metrics")
			cw, err := fetchWeather(ctx, opts.lat, opts.lng)
			if err != nil {
				logger.Error("fetching weather failed", "err", err)
				continue
			}
			if err := setMetrics(opts.locationName, *cw); err != nil {
				logger.Error("updating metrics failed", "err", err)
				continue
			}
			logger.Info("metrics fetched", "obvs-time", cw.Current.Time)
		}
	}()

	logger.Info("listening", "addr", opts.addr)

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(opts.addr, mux); err != nil && err != http.ErrServerClosed {
		logger.Error("server error", "err", err)
		os.Exit(1)
	}
	os.Exit(0)
}
