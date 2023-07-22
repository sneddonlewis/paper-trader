package model

type Position struct {
	Ticker   string  `json:"ticker"`
	Price    float64 `json:"price"`
	Quantity float64 `json:"quantity"`
}
