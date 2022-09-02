package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータを取得する(id)
	params := httprouter.ParamsFromContext(r.Context())

	// intにパース
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		// errors.Newで独自エラー作成
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	movie, err := app.models.DB.Get(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// http.StatusOKはステータスコード
	err = app.writeJSON(w, http.StatusOK, movie, "movie")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.models.DB.All()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, movies, "movies")
	if err != nil {
		app.errorJSON(w, err)
		return
	}

}

// func (app *application) deleteMovie(w http.ResponseWriter, r *http.Request) {

// }

// func (app *application) insertMovie(w http.ResponseWriter, r *http.Request) {

// }

// func (app *application) updatetMovie(w http.ResponseWriter, r *http.Request) {

// }

// func (app *application) searchMovies(w http.ResponseWriter, r *http.Request) {

// }
