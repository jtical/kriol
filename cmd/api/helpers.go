//Filename: cmd/api/helpers.go

package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// we create our method
func (app *application) readIDParam(r *http.Request) (int64, error) {
	//use the "paramsfromcontect()" function to get the request context on a slice
	params := httprouter.ParamsFromContext(r.Context())
	//get the value of the id form the parameter
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}
