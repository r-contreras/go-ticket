package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func (app *Application) StatusHandler(w http.ResponseWriter, r *http.Request) {
	currentStatus := AppStatus{
		Status:      "Available",
		Environment: app.Config.Env,
		Version:     "1.0.0",
	}

	js, err := json.MarshalIndent(currentStatus, "", "\t")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
