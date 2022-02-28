package api

import (
	utils "backend/cmd/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) GetMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		app.Logger.Println(errors.New("GetMovie Handler: Invalid id parameter"))
		return
	}

	movie, err := app.Model.GetMovie(id)

	if err != nil {
		app.Logger.Println(err)
	}

	err = utils.WriteJsonResponse("movie", movie, http.StatusOK, w)
	if err != nil {
		app.Logger.Println(errors.New("GetMovie Handler: Could not write JSON response"))
	}
}

func (app *Application) GetAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.Model.GetAllMovies()
	if err != nil {
		app.Logger.Println(err)
	}

	err = utils.WriteJsonResponse("movies", movies, http.StatusOK, w)
	if err != nil {
		app.Logger.Println(errors.New("GetAllMovies Handler: Could not write JSON response"))
	}
}
