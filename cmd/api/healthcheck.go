//Filename: cmd/api/healthcheck.go

package main

import (
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
	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encounter a problem and could not process your request", http.StatusInternalServerError)
		return
	}
}
