package web

import (
	"github.com/gin-gonic/gin"
	"paper-trader/db"
)

type App struct {
	d         db.DB
	resources []Resource
	router    *gin.Engine
}

type Resource interface {
	GetEndpoints() []Endpoint
}

type Endpoint struct {
	Method  string
	Route   string
	Handler gin.HandlerFunc
}
