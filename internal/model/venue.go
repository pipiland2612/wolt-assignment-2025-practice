package model

type Coordinates struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

type DistanceRange struct {
	Min int `json:"min"`
	Max int `json:"max"`
	A   int `json:"a"`
	B   int `json:"b"`
}

type Venue struct {
	Coords Coordinates `json:"coordinates"`
}

type VenueDynamic struct {
	Venue
	BasePrice      float64       `json:"base_price"`
	OrderMin       float64       `json:"order_minimum_no_surcharge"`
	DistanceRanges DistanceRange `json:"distance_ranges"`
}
