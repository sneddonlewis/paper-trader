package web

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"paper-trader/db"
)

type App struct {
	d                db.DB
	positionResource Resource
	router           chi.Router
}

type Resource interface {
	Handlers() map[string]http.HandlerFunc
}

func NewApp(d db.DB, resource Resource, cors bool) App {
	app := App{
		d:                d,
		positionResource: resource,
		router:           chi.NewRouter(),
	}
	// app should have an array of pointer to resource objects
	// each should be the interface with Handlers() map[string]http.HandleFunc
	// that way they can all be wired up here and dependency injected in server.go
	techHandler := app.GetTechnologies
	if !cors {
		app.router.Use(DisableCorsMiddleware)
	}
	for key, value := range app.positionResource.Handlers() {
		app.router.HandleFunc(key, value)
	}
	app.router.HandleFunc("/api/technologies", techHandler)

	app.router.Handle("/*", http.StripPrefix("/", http.FileServer(http.Dir("/webapp"))))
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

func (a *App) GetTechnologies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	technologies, err := a.d.GetTechnologies()
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(technologies)
	if err != nil {
		SendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func SendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}
