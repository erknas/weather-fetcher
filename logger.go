package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type logger struct {
	next WeatherFetcher
}

func NewLogger(next WeatherFetcher) WeatherFetcher {
	return &logger{
		next: next,
	}
}

func (s *logger) FetchWeather(ctx context.Context, city string) (temp float64, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"requestID":   ctx.Value("requestID"),
			"took":        time.Since(start),
			"err":         err,
			"temperature": temp,
			"city":        city,
		}).Info("weatcher fetch")
	}(time.Now())

	return s.next.FetchWeather(ctx, city)
}
