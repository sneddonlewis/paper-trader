package web

import (
	"github.com/gin-gonic/gin"
	"log"
	"paper-trader/db"
)

func NewApp(d db.DB, resources []Resource, cors bool) App {
	app := App{
		d:         d,
		resources: resources,
		router:    gin.Default(),
	}

	for _, resource := range app.resources {
		for _, endpoint := range resource.GetEndpoints() {
			app.router.Handle(endpoint.Method, endpoint.Route, endpoint.Handler)
		}
	}
	log.Println(app.router)

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

func SendErr(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}
