package model

type Request struct {
	VenueSlug  string  `json:"venue_slug"`
	CartValue  float64 `json:"cart_value"`
	UserCoords Location
}

type Response struct {
	TotalPrice     float64 `json:"total_price"`
	OrderSurcharge float64 `json:"small_order_surcharge"`
	CartValue      float64 `json:"cart_value"`
	Delivery       Delivery
}

type Delivery struct {
	Fee      float64 `json:"fee"`
	Distance float64 `json:"distance"`
}
