package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"paper-trader/db"
	"paper-trader/model"
	"strconv"
)

type PortfolioResource struct {
	portfolioRepo *db.PortfolioRepo
}

func NewPortfolioResource(portfolioRepo *db.PortfolioRepo) *PortfolioResource {
	return &PortfolioResource{portfolioRepo: portfolioRepo}
}

func (r *PortfolioResource) GetEndpoints() []Endpoint {
	return []Endpoint{
		{http.MethodGet, "/api/portfolio/:id", r.GetPortfolioByID},
	}
}

func (r *PortfolioResource) GetPortfolioByID(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.Atoi(idStr)
	if err != nil {
		SendErr(c, http.StatusBadRequest, "Invalid Portfolio ID provided")
		return
	}
	id := int32(id64)

	portfolio, err := r.portfolioRepo.GetPortfolioById(id)
	portfolioView := mapPortfolioView(portfolio)
	if err != nil {
		SendErr(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, portfolioView)
}

func mapPortfolioView(portfolio *model.Portfolio) *model.PortfolioView {
	return &model.PortfolioView{
		ID:              portfolio.ID,
		UserID:          portfolio.UserID,
		Name:            portfolio.Name,
		Value:           portfolio.Value,
		OpenPositions:   portfolio.OpenPositions,
		ClosedPositions: toClosedPositionViewSlice(portfolio.ClosedPositions),
	}
}

func toClosedPositionViewSlice(closedPositions []*model.ClosedPosition) []*model.ClosedPositionView {
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
