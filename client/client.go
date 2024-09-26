package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zeze322/weather-fetcher/types"
)

type Client struct {
	endpoint string
}

func New(endpoint string) *Client {
	return &Client{
		endpoint: endpoint,
	}
}

func (c *Client) FetchWeather(ctx context.Context, city string) (*types.WeatherResponse, error) {
	endpoint := fmt.Sprintf("%s?city=%s", c.endpoint, city)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	tempResp := new(types.WeatherResponse)
	if err := json.NewDecoder(resp.Body).Decode(&tempResp.Temp); err != nil {
		return nil, err
	}

	return tempResp, nil
}
