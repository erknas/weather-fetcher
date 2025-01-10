package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/zeze322/weather-fetcher/types"
)

type logger struct {
	next WeatherFetcher
}

type RequestIDKey struct{}

func NewLogger(next WeatherFetcher) WeatherFetcher {
	return &logger{
		next: next,
	}
}

func (s *logger) FetchWeather(ctx context.Context, city string) (resp types.WeatherResponse, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"city":      city,
			"time":      start,
			"requestID": ctx.Value(RequestIDKey{}),
			"took":      time.Since(start),
			"err":       err,
			"response":  resp,
		}).Info("fetched weather")
	}(time.Now())

	return s.next.FetchWeather(ctx, city)
}
