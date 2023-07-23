package web

import "paper-trader/model"

func MapPortfolioView(portfolio *model.Portfolio) *model.PortfolioView {
	return &model.PortfolioView{
		ID:              portfolio.ID,
		UserID:          portfolio.UserID,
		Name:            portfolio.Name,
		Value:           portfolio.Value,
		OpenPositions:   portfolio.OpenPositions,
		ClosedPositions: MapClosedPositionViewSlice(portfolio.ClosedPositions),
	}
}

func MapClosedPositionViewSlice(closedPositions []*model.ClosedPosition) []*model.ClosedPositionView {
	var closedPositionViews []*model.ClosedPositionView

	for _, closedPosition := range closedPositions {
		closedPositionView := &model.ClosedPositionView{
			ID:          closedPosition.ID,
			PortfolioID: closedPosition.PortfolioID,
			Ticker:      closedPosition.Ticker,
			Price:       closedPosition.Price,
			Quantity:    closedPosition.Quantity,
			OpenedAt:    closedPosition.OpenedAt,
			ClosedAt:    closedPosition.ClosedAt,
			ClosePrice:  closedPosition.ClosePrice.Float64,
			Profit:      closedPosition.Profit.Float64,
		}

		closedPositionViews = append(closedPositionViews, closedPositionView)
	}

	return closedPositionViews
}

func MapClosedPositionView(closedPosition *model.ClosedPosition) *model.ClosedPositionView {
	return &model.ClosedPositionView{
		ID:          closedPosition.ID,
		PortfolioID: closedPosition.PortfolioID,
		Ticker:      closedPosition.Ticker,
		Price:       closedPosition.Price,
		Quantity:    closedPosition.Quantity,
		OpenedAt:    closedPosition.OpenedAt,
		ClosedAt:    closedPosition.ClosedAt,
		ClosePrice:  closedPosition.ClosePrice.Float64,
		Profit:      closedPosition.Profit.Float64,
	}
}
