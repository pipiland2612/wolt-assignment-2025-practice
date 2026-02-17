package model

type VenueResponse struct {
	Venue Venue `json:"venue_raw"`
}

type Venue struct {
	Location      *Location      `json:"location,omitempty"`
	DeliverySpecs *DeliverySpecs `json:"delivery_specs,omitempty"`
}

type Location struct {
	Coordinate []float64 `json:"coordinates"`
}

type DeliverySpecs struct {
	OrderMin        float64          `json:"order_minimum_no_surcharge"`
	DeliveryPricing *DeliveryPricing `json:"delivery_pricing,omitempty"`
}

type DeliveryPricing struct {
	BasePrice      float64         `json:"base_price"`
	DistanceRanges []DistanceRange `json:"distance_ranges"`
}

type DistanceRange struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	A   float64 `json:"a"`
	B   float64 `json:"b"`
}
