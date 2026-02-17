package calculator

import (
	"errors"
	"fmt"
	"golang-api-practice/internal/model"
	"math"
)

func TotalFee(cartValue int, venue model.VenueDynamic) (float64, error) {
	smallOrderSurcharge := calcSmallOrderSurcharge(float64(cartValue), venue.OrderMin)
	smallOrderSurcharge = math.Max(smallOrderSurcharge, 0)

	distance := calcDistance(model.Coordinates{}, venue.Coords)
	deliveryFee, err := calcDeliveryFee(distance, venue)
	if err != nil {
		return 0, fmt.Errorf("error calculating delivery fee: %w", err)
	}
	return deliveryFee + smallOrderSurcharge + float64(cartValue), nil
}

func calcSmallOrderSurcharge(cartValue, orderMin float64) float64 {
	return cartValue - orderMin
}

func calcDeliveryFee(d float64, venue model.VenueDynamic) (float64, error) {
	for _, r := range venue.DistanceRanges {

		dMin := float64(r.Min)
		dMax := float64(r.Max)

		if r.Max == 0 {
			if d < dMin {
				return calcDistanceRangeFee(float64(r.A), float64(r.B), venue.BasePrice, d), nil
			}
			continue
		}

		if d >= dMin && d < dMax {
			return calcDistanceRangeFee(float64(r.A), float64(r.B), venue.BasePrice, d), nil
		}
	}

	return 0, errors.New("venue distance is out of range")
}

func calcDistanceRangeFee(a, b, basePrice, d float64) float64 {
	return math.Round(basePrice + a + b*d/10.0)
}
