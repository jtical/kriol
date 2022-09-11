//Filename: cmd/api/routes.go

package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// create a method that returns a http router
func (app *application) routes() *httprouter.Router {
	//Create a new httrouter router instance
	router := httprouter.New()
	//implement error handling in router
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedesponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/entries", app.createEntryHandler)
	router.HandlerFunc(http.MethodGet, "/v1/entries/:id", app.showEntryHandler)

	return router
}
