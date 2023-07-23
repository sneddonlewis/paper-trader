package model

import "time"

type OpenPositionRequest struct {
	Ticker   string  `json:"ticker"`
	Quantity float64 `json:"quantity"`
}

type Position struct {
	ID       int32     `json:"id"`
	Ticker   string    `json:"ticker"`
	Price    float64   `json:"price"`
	Quantity float64   `json:"quantity"`
	OpenedAt time.Time `json:"opened-at"`
}

type ClosedPosition struct {
	ID         int32     `json:"id"`
	Ticker     string    `json:"ticker"`
	Price      float64   `json:"price"`
	Quantity   float64   `json:"quantity"`
	OpenedAt   time.Time `json:"opened-at"`
	ClosePrice float64   `json:"close-price"`
	ClosedAt   time.Time `json:"closed-at"`
}
