package calculator

import (
	"errors"
	"fmt"
	"golang-api-practice/internal/model"
	"math"
)

func TotalFee(cartValue float64, userCoord *model.Location, venue *model.Venue) (*model.Response, error) {
	smallOrderSurcharge := calcSmallOrderSurcharge(cartValue, venue.DeliverySpecs.OrderMin)
	smallOrderSurcharge = math.Max(smallOrderSurcharge, 0.0)

	distance := calcDistance(userCoord, venue.Location)
	deliveryFee, err := calcDeliveryFee(distance, venue)
	if err != nil {
		return &model.Response{}, fmt.Errorf("error calculating delivery fee: %w", err)
	}
	return &model.Response{
		TotalPrice:     smallOrderSurcharge + deliveryFee + cartValue,
		OrderSurcharge: smallOrderSurcharge,
		CartValue:      cartValue,
		Delivery: model.Delivery{
			Fee:      deliveryFee,
			Distance: distance,
		},
	}, nil
}

func calcSmallOrderSurcharge(cartValue, orderMin float64) float64 {
	return cartValue - orderMin
}

func calcDeliveryFee(d float64, venue *model.Venue) (float64, error) {
	for _, r := range venue.DeliverySpecs.DeliveryPricing.DistanceRanges {
		basePrice := venue.DeliverySpecs.DeliveryPricing.BasePrice

		if r.Max == 0 {
			if d < r.Min {
				return calcDistanceRangeFee(r.A, r.B, basePrice, d), nil
			}
			continue
		}

		if d >= r.Min && d < r.Max {
			return calcDistanceRangeFee(r.A, r.B, basePrice, d), nil
		}
	}

	return 0, errors.New("venue distance is out of range")
}

func calcDistanceRangeFee(a, b, basePrice, d float64) float64 {
	return math.Round(basePrice + a + b*d/10.0)
}
