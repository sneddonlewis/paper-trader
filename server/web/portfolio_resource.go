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
	if err != nil {
		SendErr(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, portfolio)
}

func toClosedPositionResponse(closedPosition *model.ClosedPosition) *model.ClosedPositionResponse {
	closedPositionResponse := &model.ClosedPositionResponse{
		ID:          closedPosition.ID,
		PortfolioID: closedPosition.PortfolioID,
		Ticker:      closedPosition.Ticker,
		Price:       closedPosition.Price,
		Quantity:    closedPosition.Quantity,
		OpenedAt:    closedPosition.OpenedAt,
		ClosedAt:    closedPosition.ClosedAt,
	}

	if closedPosition.ClosePrice.Valid {
		closedPositionResponse.ClosePrice = closedPosition.ClosePrice.Float64
	}

	if closedPosition.Profit.Valid {
		closedPositionResponse.Profit = closedPosition.Profit.Float64
	}

	return closedPositionResponse
}
