package handler

import (
	"encoding/json"
	"fmt"
	"golang-api-practice/internal/model"
	"golang-api-practice/internal/service"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const (
	_venueSlug = "venue_slug"
	_cartValue = "cart_value"
	_userLat   = "user_lat"
	_userLon   = "user_lon"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) DeliveryPrice(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req, err := parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.service.CalculateTotalFee(ctx, req)
	if err != nil {
		log.Printf("CalculateTotalFee error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func parseRequest(r *http.Request) (*model.Request, error) {
	q := r.URL.Query()

	venue := q.Get(_venueSlug)
	if venue == "" {
		return &model.Request{}, fmt.Errorf("missing venue_slug")
	}

	cartValue, err := parseFloatParam(q, _cartValue)
	if err != nil {
		return &model.Request{}, fmt.Errorf("invalid cart_value")
	}

	lat, err := parseFloatParam(q, _userLat)
	if err != nil {
		return &model.Request{}, fmt.Errorf("invalid user_lat")
	}

	lon, err := parseFloatParam(q, _userLon)
	if err != nil {
		return &model.Request{}, fmt.Errorf("invalid user_lon")
	}

	return &model.Request{
		VenueSlug: venue,
		CartValue: cartValue,
		UserCoords: model.Location{
			Coordinate: []float64{lon, lat},
		},
	}, nil
}

func parseFloatParam(q url.Values, key string) (float64, error) {
	value := q.Get(key)
	if value == "" {
		return 0, fmt.Errorf("missing %s", key)
	}
	return strconv.ParseFloat(value, 64)
}
