package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/zeze322/weather-fetcher/types"
)

const (
	apiKey  = "afcca19db4b548c38c1171000251001"
	baseURL = "http://api.weatherapi.com/v1/current.json?"
)

type WeatherFetcher interface {
	FetchWeather(context.Context, string) (types.WeatherResponse, error)
}

type weatherFetcher struct{}

type Response struct {
	types.WeatherResponse
	err error
}

func (s *weatherFetcher) FetchWeather(ctx context.Context, city string) (types.WeatherResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*500)
	defer cancel()

	respch := make(chan Response)

	go func() {
		weather, err := fetchWeather(city)
		respch <- Response{
			WeatherResponse: weather,
			err:             err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return types.WeatherResponse{}, fmt.Errorf("context canceled")
		case resp := <-respch:
			return resp.WeatherResponse, resp.err
		}
	}
}

func fetchWeather(city string) (types.WeatherResponse, error) {
	URL := fmt.Sprintf("%skey=%s&q=%s&aqi=no", baseURL, apiKey, city)

	resp, err := http.Get(URL)
	if err != nil {
		return types.WeatherResponse{}, err
	}
	defer resp.Body.Close()

	var weather types.WeatherResponse

	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return types.WeatherResponse{}, err
	}

	return weather, nil
}
