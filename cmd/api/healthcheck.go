//Filename: cmd/api/healthcheck.go

package main

import (
	"encoding/json"
	"net/http"
)

// create a healthcheck handler
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	//create a map to hold healthcheck data
	data := map[string]string{
		"status":      "available",
		"enviornment": app.config.env,
		"version":     version,
	}
	//convert our map into a json object
	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encounter a problem and could not process your request", http.StatusInternalServerError)
		return
	}
	//a new line to byte slices to make viewing on the terminal easier
	js = append(js, '\n')
	// specify that we will server our responses using JSON
	w.Header().Set("Content-Type", "application/json")
	//we write out the []byte slices conating the json response boday
	w.Write(js)

}
