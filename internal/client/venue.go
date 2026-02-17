package client

import (
	"encoding/json"
	"fmt"
	"golang-api-practice/internal/model"
	"net/http"
)

const _apiPrefix = "https://consumer-api.development.dev.woltapi.com/home-assignment-api/v1/venues/"

type result struct {
	venue model.VenueResponse
	err   error
}

func FetchApi(venue string) (model.Venue, error) {
	staticChan := make(chan result)
	dynamicChan := make(chan result)

	go func() {
		res, err := fetchStatic(venue)
		staticChan <- result{res, err}
	}()

	go func() {
		res, err := fetchDynamic(venue)
		dynamicChan <- result{res, err}
	}()

	staticResult := <-staticChan
	dynamicResult := <-dynamicChan

	if staticResult.err != nil {
		return model.Venue{}, fmt.Errorf("static endpoint fetch failed: %w", staticResult.err)
	}

	if dynamicResult.err != nil {
		return model.Venue{}, fmt.Errorf("dynamic endpoint fetch failed: %w", dynamicResult.err)
	}

	return model.Venue{
		Location:      staticResult.venue.Venue.Location,
		DeliverySpecs: dynamicResult.venue.Venue.DeliverySpecs,
	}, nil
}

func fetchDynamic(venue string) (model.VenueResponse, error) {
	url := fmt.Sprintf("%v%v/dynamic", _apiPrefix, venue)
	return fetchAndParseURL(url)
}

func fetchStatic(venue string) (model.VenueResponse, error) {
	url := fmt.Sprintf("%v%v/static", _apiPrefix, venue)
	return fetchAndParseURL(url)
}

func fetchAndParseURL(url string) (model.VenueResponse, error) {
	client := &http.Client{}

	resp, err := client.Get(url)
	if err != nil {
		return model.VenueResponse{}, err
	}
	defer resp.Body.Close()

	var result model.VenueResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return model.VenueResponse{}, err
	}

	return result, nil
}
