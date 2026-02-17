package service

import (
	"context"
	"golang-api-practice/internal/calculator"
	"golang-api-practice/internal/client"
	"golang-api-practice/internal/model"
)

type service struct {
	client *client.ApiClient
}

type Service interface {
	CalculateTotalFee(ctx context.Context, req *model.Request) (*model.Response, error)
}

func NewService(client *client.ApiClient) Service {
	return &service{client: client}
}

func (s *service) CalculateTotalFee(ctx context.Context, req *model.Request) (*model.Response, error) {
	venue, err := s.client.FetchApi(ctx, req.VenueSlug)
	if err != nil {
		return &model.Response{}, err
	}

	return calculator.TotalFee(req.CartValue, &req.UserCoords, venue)
}
