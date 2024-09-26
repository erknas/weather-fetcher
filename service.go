package main

import (
	"context"
	"fmt"
)

type WeatherFetcher interface {
	FetchWeather(context.Context, string) (float64, error)
}

type weatherFetcher struct{}

var weatherMock = map[string]float64{
	"Moscow":           18.1,
	"Omsk":             9.0,
	"Saint-Petersburg": 21.5,
}

func (s *weatherFetcher) FetchWeather(ctx context.Context, city string) (float64, error) {
	temp, ok := weatherMock[city]
	if !ok {
		return temp, fmt.Errorf("%s doesn't exist", city)
	}

	return temp, nil
}
