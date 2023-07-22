package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"paper-trader/db"
	"paper-trader/service"
	"paper-trader/web"
)

func main() {
	d, err := sql.Open("mysql", dataSource())
	if err != nil {
		log.Fatal(err)
	}
	defer d.Close()
	// CORS is enabled only in prod profile
	cors := os.Getenv("profile") == "prod"

	positionRepo := db.NewPositionRepo(d)
	tradeService := service.NewTradeService(positionRepo, service.PricingService{})
	positionResource := web.NewPositionResource(tradeService)
	resources := []web.Resource{
		positionResource,
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
	return "goxygen:" + pass + "@tcp(" + host + ":3306)/goxygen"
}
