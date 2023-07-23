package model

import (
	"database/sql"
	"time"
)

type OpenPositionRequest struct {
	PortfolioID int32   `json:"portfolio_id"`
	Ticker      string  `json:"ticker"`
	Quantity    float64 `json:"quantity"`
}

type Position struct {
	ID          int32      `json:"id"`
	PortfolioID int32      `json:"portfolio_id"`
	Ticker      string     `json:"ticker"`
	Price       float64    `json:"price"`
	Quantity    float64    `json:"quantity"`
	OpenedAt    *time.Time `json:"opened-at"`
}

type ClosedPosition struct {
	ID          int32           `json:"id"`
	PortfolioID int32           `json:"portfolio_id"`
	Ticker      string          `json:"ticker"`
	Price       float64         `json:"price"`
	Quantity    float64         `json:"quantity"`
	OpenedAt    *time.Time      `json:"opened-at"`
	ClosePrice  sql.NullFloat64 `json:"close-price"`
	ClosedAt    *time.Time      `json:"closed-at"`
	Profit      sql.NullFloat64 `json:"profit"`
}
type ClosedPositionResponse struct {
	ID          int32      `json:"id"`
	PortfolioID int32      `json:"portfolio_id"`
	Ticker      string     `json:"ticker"`
	Price       float64    `json:"price"`
	Quantity    float64    `json:"quantity"`
	OpenedAt    *time.Time `json:"opened-at"`
	ClosePrice  float64    `json:"close-price"`
	ClosedAt    *time.Time `json:"closed-at"`
	Profit      float64    `json:"profit"`
}
