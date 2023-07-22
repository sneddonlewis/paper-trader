package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"paper-trader/db"
	"paper-trader/model"
	"strings"
)

type PositionResource struct {
	repo db.PositionRepo
}

func NewPositionResource(repo db.PositionRepo) *PositionResource {
	return &PositionResource{repo: repo}
}

func (r *PositionResource) GetEndpoints() []Endpoint {
	return []Endpoint{
		{http.MethodGet, "/api/positions", r.GetPositions},
		{http.MethodPost, "/api/position/:ticker/close", r.ClosePosition},
	}
}

func (r *PositionResource) GetPositions(c *gin.Context) {
	positions, err := r.repo.GetPositions()
	if err != nil {
		SendErr(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, positions)
}

func (r *PositionResource) ClosePosition(c *gin.Context) {
	ticker := strings.ToUpper(c.Param("ticker"))
	positions, err := r.repo.GetPositions()
	if err != nil {
		SendErr(c, http.StatusInternalServerError, err.Error())
		return
	}
	positions = withTicker(positions, ticker)
	exp := exposure(positions)
	closingPosition := r.openPosition(model.Position{
		Ticker:   ticker,
		Price:    100.0,
		Quantity: exp * -1,
	})
	c.JSON(http.StatusOK, closingPosition)
}

func (r *PositionResource) openPosition(position model.Position) model.Position {
	return position
}

func exposure(positions []*model.Position) float64 {
	amount := 0.0
	for _, position := range positions {
		amount = amount + position.Quantity
	}
	return amount
}

func withTicker(positions []*model.Position, ticker string) []*model.Position {
	var result []*model.Position
	for _, position := range positions {
		if position.Ticker == ticker {
			result = append(result, position)
		}
	}
	return result
}
