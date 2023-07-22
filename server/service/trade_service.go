package service

import (
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
	return s.positionRepo.GetPositions()
}
