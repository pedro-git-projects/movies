package main

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/pedro-git-projects/movies/models"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Println(errors.New("invalid id parameter"))
	}

	app.logger.Println("id is", id)

	movie := models.Movie{
		ID:          id, // gotten by the parameters
		Title:       "place holder",
		Description: "some description",
		Year:        20201,
		ReleaseDate: time.Date(2021, 01, 01, 01, 01, 0, 0, time.Local),
		Runtime:     100,
		Rating:      5,
		MPAARating:  "PG-13",
		CreatedAt:   time.Now(),
		UdateddAt:   time.Now(),
	}
	err = app.writeJSON(w, http.StatusOK, movie, "movie")
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {
}
