package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	// getメソッド
	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	// 第三引数にハンドラーを渡す
	router.HandlerFunc(http.MethodGet, "/v1/movie/:id", app.getOneMovie)
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.getAllMovies)

	router.HandlerFunc(http.MethodGet, "/v1/genres", app.getAllGenres)

	return app.enableCORS(router)
}
