package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zeze322/weather-fetcher/types"
)

const (
	apiKey  = "29ef695437cc4131b78140816240310"
	baseURL = "http://api.weatherapi.com/v1/current.json?"
)

type WeatherFetcher interface {
	FetchWeather(context.Context, string) (*types.WeatherResponse, error)
}

type weatherFetcher struct{}

func (s *weatherFetcher) FetchWeather(ctx context.Context, city string) (*types.WeatherResponse, error) {
	URL := fmt.Sprintf("%skey=%s&q=%s&aqi=no", baseURL, apiKey, city)

	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}

	weather := new(types.WeatherResponse)

	if err := json.NewDecoder(resp.Body).Decode(weather); err != nil {
		return nil, err
	}

	return weather, nil
}
