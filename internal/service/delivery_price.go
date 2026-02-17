package service

import (
	"context"
	"golang-api-practice/internal/calculator"
	"golang-api-practice/internal/client"
	"golang-api-practice/internal/model"
)

type Service struct {
	client *client.ApiClient
}

func NewService(client *client.ApiClient) *Service {
	return &Service{client: client}
}

func (s *Service) CalculateTotalFee(ctx context.Context, req *model.Request) (*model.Response, error) {
	venue, err := s.client.FetchApi(ctx, req.VenueSlug)
	if err != nil {
		return &model.Response{}, err
	}

	return calculator.TotalFee(req.CartValue, &req.UserCoords, venue)
}
