//Filename: cmd/api/helpers.go

package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// we create our method to read id
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

// we create our method write json to create responses
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	//a new line to byte slices to make viewing on the terminal easier
	js = append(js, '\n')
	//add the headers
	for key, value := range headers {
		w.Header()[key] = value

	}
	// specify that we will server our responses using JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	//we write out the []byte slices conating the json response body
	w.Write(js)
	return nil
}
