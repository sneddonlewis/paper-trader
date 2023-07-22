package model

type OpenPositionRequest struct {
	Ticker   string  `json:"ticker"`
	Quantity float64 `json:"quantity"`
}

type Position struct {
	ID       int32   `json:"id"`
	Ticker   string  `json:"ticker"`
	Price    float64 `json:"price"`
	Quantity float64 `json:"quantity"`
}

type ClosedPosition struct {
	ID         int32   `json:"id"`
	Ticker     string  `json:"ticker"`
	Price      float64 `json:"price"`
	Quantity   float64 `json:"quantity"`
	ClosePrice float64 `json:"close-price"`
}
