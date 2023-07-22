package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"paper-trader/db"
)

type PositionResource struct {
	repo db.PositionRepo
}

func NewPositionResource(repo db.PositionRepo) *PositionResource {
	return &PositionResource{repo: repo}
}

func (resource *PositionResource) GetEndpoints() []Endpoint {
	return []Endpoint{
		{http.MethodGet, "/api/positions", resource.GetPositions},
	}
}

func (resource *PositionResource) Handlers() map[string]gin.HandlerFunc {
	return map[string]gin.HandlerFunc{
		"/api/positions": resource.GetPositions,
	}
}

func (resource *PositionResource) GetPositions(c *gin.Context) {
	positions, err := resource.repo.GetPositions()
	if err != nil {
		SendErr(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, positions)
}
