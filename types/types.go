package types

type WeatherResponse struct {
	Location Location
	Current  Current
}

type Location struct {
	Name      string  `json:"name"`
	Region    string  `json:"region"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	LocalTime string  `json:"localtime"`
}

type Current struct {
	Temperature float64 `json:"temp_c"`
	Wind        float64 `json:"wind_kph"`
	Pressure    float64 `json:"pressure_mb"`
	Humidity    float64 `json:"humidity"`
	Feelslike   float64 `json:"feelslike_c"`
	Condition   Condition
}

type Condition struct {
	Name string `json:"text"`
}
