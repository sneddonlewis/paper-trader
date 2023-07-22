package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"paper-trader/model"
	"paper-trader/service"
	"strconv"
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
		{http.MethodPost, "/api/position/:id/close", r.ClosePosition},
		{http.MethodPost, "/api/position", r.OpenPosition},
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

func (r *PositionResource) OpenPosition(c *gin.Context) {
	var requestModel model.OpenPositionRequest
	if err := c.BindJSON(&requestModel); err != nil {
		SendErr(c, http.StatusBadRequest, "Invalid request body")
		return
	}
	opened, err := r.tradeService.OpenPosition(&requestModel)
	if err != nil {
		SendErr(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, opened)
}

func (r *PositionResource) ClosePosition(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.Atoi(idStr)
	if err != nil {
		SendErr(c, http.StatusBadRequest, "Invalid ID provided")
		return
	}
	id := int32(id64)
	closingPosition, err := r.tradeService.ClosePosition(id)
	if err != nil {
		SendErr(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, closingPosition)
}
