package model

type Position struct {
	Ticker    string  `json:"ticker"`
	Direction string  `json:"direction"`
	Price     float64 `json:"price"`
	Quantity  float64 `json:"quantity"`
}
