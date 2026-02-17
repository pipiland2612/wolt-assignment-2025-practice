package service

import (
	"context"
	"golang-api-practice/internal/calculator"
	"golang-api-practice/internal/client"
	"golang-api-practice/internal/model"
)

func CalculateTotalFee(ctx context.Context, req *model.Request) (*model.Response, error) {
	venue, err := client.FetchApi(ctx, req.VenueSlug)
	if err != nil {
		return &model.Response{}, err
	}

	return calculator.TotalFee(req.CartValue, &req.UserCoords, venue)
}
