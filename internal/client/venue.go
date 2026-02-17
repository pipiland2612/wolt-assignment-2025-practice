package client

import (
	"encoding/json"
	"fmt"
	"golang-api-practice/internal/model"
	"net/http"
)

const _apiPrefix = "https://consumer-api.development.dev.woltapi.com/home-assignment-api/v1/venues/"

func FetchApi(venue string) (model.Venue, error) {
	static, err := fetchStatic(venue)
	if err != nil {
		return model.Venue{}, err
	}

	dynamic, err := fetchDynamic(venue)
	if err != nil {
		return model.Venue{}, err
	}

	return model.Venue{
		Location:      static.Venue.Location,
		DeliverySpecs: dynamic.Venue.DeliverySpecs,
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
