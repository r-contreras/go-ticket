package implementations

import (
	api "backend/cmd/api"
	interfaces "backend/cmd/interfaces"
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Application struct{
	Config api.Config
	Logger *log.Logger
	interfaces.IApplication
}

func (app *Application) Routes() *httprouter.Router {
	router := httprouter.New()
	
	//setup all routes
	router.HandlerFunc(http.MethodGet, "/status", app.StatusHandler)

	return router
}

func (app *Application) StatusHandler(w http.ResponseWriter, r *http.Request) {
	currentStatus := api.AppStatus {
		Status: "Available",
		Environment: app.Config.Env,
		Version: "1.0.0",
	}

	js, err := json.MarshalIndent(currentStatus, "", "\t")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}