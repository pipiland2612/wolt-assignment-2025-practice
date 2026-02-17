package client

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-api-practice/internal/model"
	"net/http"
)

const _apiPrefix = "https://consumer-api.development.dev.woltapi.com/home-assignment-api/v1/venues/"

type result struct {
	venue model.Venue
	err   error
}

func FetchApi(ctx context.Context, venue string) (model.Venue, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	staticChan := make(chan result, 1)
	dynamicChan := make(chan result, 1)

	go fetchHelper(ctx, staticChan, venue, fetchStatic)
	go fetchHelper(ctx, dynamicChan, venue, fetchDynamic)

	var static, dynamic result

	for i := 0; i < 2; i++ {
		select {
		case s := <-staticChan:
			static = s
			if s.err != nil {
				return model.Venue{}, fmt.Errorf("static endpoint fetch failed: %w", s.err)
			}
		case d := <-dynamicChan:
			dynamic = d
			if d.err != nil {
				return model.Venue{}, fmt.Errorf("dynamic endpoint fetch failed: %w", d.err)
			}
		case <-ctx.Done():
			return model.Venue{}, ctx.Err()
		}
	}

	return model.Venue{
		Location:      static.venue.Location,
		DeliverySpecs: dynamic.venue.DeliverySpecs,
	}, nil
}

func fetchHelper(ctx context.Context, resultChan chan<- result, venue string, fn func(ctx context.Context, venue string) (model.Venue, error)) {
	res, err := fn(ctx, venue)
	select {
	case resultChan <- result{res, err}:
	case <-ctx.Done():
		return
	}
}

func fetchDynamic(ctx context.Context, venue string) (model.Venue, error) {
	url := fmt.Sprintf("%v%v/dynamic", _apiPrefix, venue)
	return fetchAndParseURL(ctx, url)
}

func fetchStatic(ctx context.Context, venue string) (model.Venue, error) {
	url := fmt.Sprintf("%v%v/static", _apiPrefix, venue)
	return fetchAndParseURL(ctx, url)
}

func fetchAndParseURL(ctx context.Context, url string) (model.Venue, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return model.Venue{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() != nil {
			return model.Venue{}, ctx.Err()
		}
		return model.Venue{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Venue{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var res model.VenueResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return model.Venue{}, err
	}

	return res.Venue, nil
}
