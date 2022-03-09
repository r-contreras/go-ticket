package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) GetMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		app.Logger.Println("GetMovie Handler:", err)
		app.WriteBadRequest(w)
		return
	}

	movie, err := app.Model.GetMovie(id)

	if err != nil {
		app.Logger.Println("GetMovie Handler:", err)
	}

	err = app.WriteJsonResponse("movie", movie, http.StatusOK, w)

	if err != nil {
		app.Logger.Println("GetMovie Handler:", err)
	}
}

func (app *Application) GetAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.Model.GetAllMovies()
	if err != nil {
		app.Logger.Println(err)
	}

	err = app.WriteJsonResponse("movies", movies, http.StatusOK, w)
	if err != nil {
		app.Logger.Println("GetAllMovies Handler:", err)
	}
}
