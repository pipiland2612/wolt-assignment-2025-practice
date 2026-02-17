package calculator

import (
	"golang-api-practice/internal/model"
	"math"
)

const EarthRadiusM = 6371000.0

// calcDistance use Haversine formula to calculate straight line distance between user and venue
func calcDistance(userCoords, venueCoords *model.Location) float64 {
	userLon := userCoords.Coordinate[0]
	userLat := userCoords.Coordinate[1]

	venueLon := venueCoords.Coordinate[0]
	venueLat := venueCoords.Coordinate[1]

	userLatRad := toRadian(userLat)
	venueLatRad := toRadian(venueLat)
	deltaLat := toRadian(venueLat - userLat)
	deltaLon := toRadian(venueLon - userLon)

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(userLatRad)*math.Cos(venueLatRad)*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return math.Round(EarthRadiusM * c)
}

func toRadian(value float64) float64 {
	return value * math.Pi / 180.0
}
