//Filename: cmd/api/entries.go

package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// createEntryHandler for the POST /v1/entries endpoint
func (app *application) createEntryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "creating a new entry..")
}

// showEntryHandler for the GET /v1/entries/:id endpoint
func (app *application) showEntryHandler(w http.ResponseWriter, r *http.Request) {
	//use the "paramsfromcontect()" function to get the request context on a slice
	params := httprouter.ParamsFromContext(r.Context())
	//get the value of the id form the parameter
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	//display the school id
	fmt.Fprintf(w, "show the details for entry %d\n", id)

}
