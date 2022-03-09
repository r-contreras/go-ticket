package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) GetGenre(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		app.Logger.Println("GetGenre Handler: ", err)
		app.WriteBadRequest(w)
		return
	}

	genre, err := app.Model.GetGenre(id)

	if err != nil {
		app.Logger.Println("GetGenre Handler: ", err)
		return
	}

	err = app.WriteJsonResponse("genre", genre, http.StatusOK, w)

	if err != nil {
		app.Logger.Println("GetGenre Handler: ", err)
	}
}

func (app *Application) GetAllGenres(w http.ResponseWriter, r *http.Request) {

	genres, err := app.Model.GetAllGenres()

	if err != nil {
		app.Logger.Println("GetAllGenres Handler: ", err)
		return
	}

	err = app.WriteJsonResponse("genres", genres, http.StatusOK, w)

	if err != nil {
		app.Logger.Println("GetAllGenres Handler: ", err)
	}
}
