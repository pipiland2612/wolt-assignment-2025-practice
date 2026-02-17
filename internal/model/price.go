package model

type Request struct {
	VenueSlug  string `json:"venue_slug"`
	CartValue  int    `json:"cart_value"`
	UserCoords Coordinates
}

type Response struct {
	TotalPrice     int `json:"total_price"`
	OrderSurcharge int `json:"small_order_surcharge"`
	CartValue      int `json:"cart_value"`
	Delivery       Delivery
}

type Delivery struct {
	Fee      int `json:"fee"`
	Distance int `json:"distance"`
}
