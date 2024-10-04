package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zeze322/weather-fetcher/types"
)

type Client struct {
	endpoint string
}

func JSONClient(endpoint string) *Client {
	return &Client{
		endpoint: endpoint,
	}
}

func (c *Client) FetchWeather(city string) (*types.WeatherResponse, error) {
	endpoint := fmt.Sprintf("%s?city=%s", c.endpoint, city)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		httpErr := map[string]any{}
		if err := json.NewDecoder(resp.Body).Decode(&httpErr); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("service responded with not status ok code: %s", httpErr["error"])
	}

	weatherResp := new(types.WeatherResponse)

	if err := json.NewDecoder(resp.Body).Decode(weatherResp); err != nil {
		return nil, err
	}

	return weatherResp, nil
}
