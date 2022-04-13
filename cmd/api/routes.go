package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodGet, "/api/v1/airports", app.getAllAirports)
	router.HandlerFunc(http.MethodGet, "/api/v1/aircraft", app.getAllAircraft)

	return router
}
