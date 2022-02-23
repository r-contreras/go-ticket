package api

import (
	"backend/cmd/models"
	utils "backend/cmd/utils"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) GetMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		app.Logger.Println(errors.New("GetMovie Handler: Invalid id parameter"))
		return
	}

	movie := models.Movie{
		Id:          id,
		Title:       "The Shawshank Redemption",
		Description: "The description",
		Year:        "1998",
		ReleaseDate: time.Date(2021, 01, 01, 01, 0, 0, 0, time.Local),
		Runtime:     100,
		Rating:      4.5,
		MPAARating:  "PG-13",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = utils.WriteJsonResponse("movie", movie, http.StatusOK, w)
	if err != nil {
		app.Logger.Println(errors.New("GetMovie Handler: Could not write JSON response"))
		return
	}
}

func (app *Application) GetAllMovies(w http.ResponseWriter, r *http.Request) {

}
