package service

import (
	"fmt"
	"paper-trader/db"
	"paper-trader/model"
)

type TradeService struct {
	positionRepo db.PositionRepo
	pricing      PricingService
}

func NewTradeService(positionRepo db.PositionRepo, pricing PricingService) *TradeService {
	return &TradeService{
		positionRepo: positionRepo,
		pricing:      pricing,
	}
}

func (s *TradeService) AllPositions() ([]*model.Position, error) {
	return s.positionRepo.GetOpenPositions()
}

func (s *TradeService) ClosePosition(id int32) (*model.ClosedPosition, error) {
	p, err := s.positionRepo.GetPositionById(id)
	if err != nil {
		return nil, err
	}
	price := s.pricing.GetSimplePrice(p.Ticker)
	position, err := s.positionRepo.ClosePosition(id, price)
	if err != nil {
		return nil, err
	}
	return position, nil
}

func (s *TradeService) ClosePositions(ticker string) (*model.Position, error) {
	positions, err := s.positionRepo.GetPositionsByTicker(ticker)
	if err != nil {
		return nil, err
	}
	exp := exposure(positions)
	if exp == 0.0 {
		return nil, fmt.Errorf("no open positions to close for ticker %v", ticker)
	}
	createModel := model.Position{
		Ticker:   ticker,
		Price:    s.pricing.GetSimplePrice(ticker),
		Quantity: exp * -1,
	}
	closingPosition, _ := s.positionRepo.OpenPosition(&createModel)
	return closingPosition, nil
}

func (s *TradeService) openPosition(position model.Position) model.Position {
	return position
}

func exposure(positions []*model.Position) float64 {
	amount := 0.0
	for _, position := range positions {
		amount = amount + position.Quantity
	}
	return amount
}
