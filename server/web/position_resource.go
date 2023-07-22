package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"paper-trader/service"
	"strings"
)

type PositionResource struct {
	tradeService *service.TradeService
}

func NewPositionResource(tradeService *service.TradeService) *PositionResource {
	return &PositionResource{tradeService: tradeService}
}

func (r *PositionResource) GetEndpoints() []Endpoint {
	return []Endpoint{
		{http.MethodGet, "/api/positions", r.GetPositions},
		{http.MethodPost, "/api/position/:ticker/close", r.ClosePosition},
	}
}

func (r *PositionResource) GetPositions(c *gin.Context) {
	positions, err := r.tradeService.AllPositions()
	if err != nil {
		SendErr(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, positions)
}

func (r *PositionResource) ClosePosition(c *gin.Context) {
	ticker := strings.ToUpper(c.Param("ticker"))
	closingPosition, err := r.tradeService.ClosePosition(ticker)
	if err != nil {
		SendErr(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, closingPosition)
}
