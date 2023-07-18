package web

import (
	"encoding/json"
	"net/http"
	"paper-trader/db"
)

type PositionResource struct {
	repo db.PositionRepo
}

func NewPositionResource(repo db.PositionRepo) *PositionResource {
	return &PositionResource{repo: repo}
}

func (resource *PositionResource) Handlers() map[string]http.HandlerFunc {
	return map[string]http.HandlerFunc{
		"/api/positions": resource.GetPositions,
	}
}

func (resource *PositionResource) GetPositions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	positions, err := resource.repo.GetPositions()
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(positions)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())
	}
}
