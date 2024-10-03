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

type requestIDKey struct{}

func NewLogger(next WeatherFetcher) WeatherFetcher {
	return &logger{
		next: next,
	}
}

func (s *logger) FetchWeather(ctx context.Context, city string) (resp *types.WeatherResponse, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"time":      start,
			"requestID": ctx.Value(requestIDKey{}),
			"took":      time.Since(start),
			"err":       err,
			"response":  resp,
			"city":      city,
		}).Info("weather fetch")
	}(time.Now())

	return s.next.FetchWeather(ctx, city)
}
