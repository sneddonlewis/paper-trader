package web

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"paper-trader/db"
)

type App struct {
	d                db.DB
	positionResource Resource
	router           *gin.Engine
}

type Resource interface {
	GetEndpoints() []Endpoint
}

type Endpoint struct {
	Method  string
	Route   string
	Handler gin.HandlerFunc
}

func NewApp(d db.DB, resource Resource, cors bool) App {
	app := App{
		d:                d,
		positionResource: resource,
		router:           gin.Default(),
	}
	// app should have an array of pointer to resource objects
	// each should be the interface with Handlers() map[string]http.HandleFunc
	// that way they can all be wired up here and dependency injected in server.go
	techHandler := app.GetTechnologies
	if !cors {
		app.router.Use(DisableCorsMiddleware)
	}
	for _, endpoint := range app.positionResource.GetEndpoints() {
		app.router.Handle(endpoint.Method, endpoint.Route, endpoint.Handler)
	}
	app.router.Handle(http.MethodGet, "/api/technologies", techHandler)
	app.router.NoRoute(func(c *gin.Context) {
		c.File("/webapp/index.html")
	})
	return app
}

func DisableCorsMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	c.Next()
}

func (a *App) Serve() error {
	log.Println("Web server is available on port 8080")
	return a.router.Run(":8080")
}

func (a *App) GetTechnologies(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	technologies, err := a.d.GetTechnologies()
	if err != nil {
		SendErr(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, technologies)
}

func SendErr(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}
