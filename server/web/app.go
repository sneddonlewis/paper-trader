package web

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"paper-trader/db"
)

type App struct {
	d            db.DB
	positionRepo db.PositionRepo
	router       *mux.Router
}

func NewApp(d db.DB, positionRepo db.PositionRepo, cors bool) App {
	app := App{
		d:            d,
		positionRepo: positionRepo,
		router:       mux.NewRouter(),
	}
	techHandler := app.GetTechnologies
	if !cors {
		app.router.Use(DisableCorsMiddleware)
	}
	app.router.HandleFunc("/api/technologies", techHandler)
	app.router.HandleFunc("/api/positions", app.GetPositions)
	return app
}

func DisableCorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		next.ServeHTTP(w, r)
	})
}

func (a *App) Serve() error {
	log.Println("Web server is available on port 8080")
	return http.ListenAndServe(":8080", a.router)
}

func (a *App) GetPositions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	positions, err := a.positionRepo.GetPositions()
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(positions)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (a *App) GetTechnologies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	technologies, err := a.d.GetTechnologies()
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(technologies)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}
