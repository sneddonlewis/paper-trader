package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"paper-trader/db"
	"paper-trader/service"
	"paper-trader/web"
)

func main() {
	d, err := sql.Open("postgres", dataSource())
	if err != nil {
		log.Fatal(err)
	}
	defer d.Close()
	// CORS is enabled only in prod profile
	cors := os.Getenv("profile") == "prod"

	positionRepo := db.NewPositionRepo(d)
	portfolioRepo := db.NewPortfolioRepo(d)
	tradeService := service.NewTradeService(positionRepo, service.PricingService{})
	positionResource := web.NewPositionResource(tradeService)
	resources := []web.Resource{
		positionResource,
		web.NewPortfolioResource(&portfolioRepo),
	}
	app := web.NewApp(db.NewDB(d), resources, cors)
	err = app.Serve()
	log.Println("Error", err)
}
func dataSource() string {
	host := "localhost"
	pass := "pass"
	if os.Getenv("profile") == "prod" {
		host = "db"
		pass = os.Getenv("db_pass")
	}
	return "host=" + host + " user=goxygen password=" + pass + " dbname=goxygen sslmode=disable"
}
