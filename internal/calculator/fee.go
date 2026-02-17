package calculator

import (
	"errors"
	"fmt"
	"golang-api-practice/internal/model"
	"math"
)

func TotalFee(cartValue int, userCoord *model.Location, venue *model.Venue) (float64, error) {
	smallOrderSurcharge := calcSmallOrderSurcharge(float64(cartValue), venue.DeliverySpecs.OrderMin)
	smallOrderSurcharge = math.Max(smallOrderSurcharge, 0)

	distance := calcDistance(userCoord, venue.Location)
	deliveryFee, err := calcDeliveryFee(distance, venue)
	if err != nil {
		return 0, fmt.Errorf("error calculating delivery fee: %w", err)
	}
	return deliveryFee + smallOrderSurcharge + float64(cartValue), nil
}

func calcSmallOrderSurcharge(cartValue, orderMin float64) float64 {
	return cartValue - orderMin
}

func calcDeliveryFee(d float64, venue *model.Venue) (float64, error) {
	for _, r := range venue.DeliverySpecs.DeliveryPricing.DistanceRanges {

		dMin := r.Min
		dMax := r.Max
		basePrice := venue.DeliverySpecs.DeliveryPricing.BasePrice

		if r.Max == 0 {
			if d < dMin {
				return calcDistanceRangeFee(r.A, r.B, basePrice, d), nil
			}
			continue
		}

		if d >= dMin && d < dMax {
			return calcDistanceRangeFee(r.A, r.B, basePrice, d), nil
		}
	}

	return 0, errors.New("venue distance is out of range")
}

func calcDistanceRangeFee(a, b, basePrice, d float64) float64 {
	return math.Round(basePrice + a + b*d/10.0)
}
