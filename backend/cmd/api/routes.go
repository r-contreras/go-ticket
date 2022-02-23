package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) Routes() *httprouter.Router {
	router := httprouter.New()

	//setup all routes
	router.HandlerFunc(http.MethodGet, "/status", app.StatusHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movie/:id", app.GetMovie)
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.GetAllMovies)

	return router
}
