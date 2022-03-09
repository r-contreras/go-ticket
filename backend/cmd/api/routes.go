package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) Routes() http.Handler {
	router := httprouter.New()

	//setup all routes
	router.HandlerFunc(http.MethodGet, "/status", app.StatusHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movie/:id", app.GetMovie)
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.GetAllMovies)
	router.HandlerFunc(http.MethodGet, "/v1/genre/:id", app.GetGenre)
	router.HandlerFunc(http.MethodGet, "/v1/genres", app.GetAllGenres)
	return app.enableCORS(router)
}
