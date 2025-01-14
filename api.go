package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type JSONAPIServer struct {
	port  string
	svc   WeatherFetcher
	cache Cacher
}

func NewJSONAPIServer(port string, svc WeatherFetcher, cache Cacher) *JSONAPIServer {
	return &JSONAPIServer{
		port:  port,
		svc:   svc,
		cache: cache,
	}
}

func (s *JSONAPIServer) Run() {
	http.HandleFunc("/", makeHTTPHandler(s.handleFetchWeather))

	log.Fatal(http.ListenAndServe(s.port, nil))
}

func (s *JSONAPIServer) handleFetchWeather(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	city := r.URL.Query().Get("city")

	val, ok := s.cache.Get(ctx, city)
	if !ok {
		weather, err := s.svc.FetchWeather(ctx, city)
		if err != nil {
			return err
		}

		if err := s.cache.Set(ctx, city, weather); err != nil {
			return err
		}

		return writeJSON(w, http.StatusOK, weather)
	}

	return writeJSON(w, http.StatusOK, val)
}

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

func makeHTTPHandler(fn APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, RequestIDKey{}, uuid.New().String())
		if err := fn(ctx, w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
	}
}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(s)

	return json.NewEncoder(w).Encode(v)
}
