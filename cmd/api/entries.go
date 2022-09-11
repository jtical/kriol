//Filename: cmd/api/entries.go

package main

import (
	"fmt"
	"net/http"
	"time"

	"kriol.joelical.net/internal/data"
)

// createEntryHandler for the POST /v1/entries endpoint
func (app *application) createEntryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "creating a new entry..")
}

// showEntryHandler for the GET /v1/entries/:id endpoint
func (app *application) showEntryHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	//create a new instance of the entries struct containting the id we ectracted from our ur and sample data
	entries := data.Entries{
		ID:        id,
		English:   "Water",
		Kriol:     "Waata",
		CreatedAt: time.Now(),
		Version:   1,
	}
	//call json so it can display in json format
	err = app.writeJSON(w, http.StatusOK, entries, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encounter a problem and could not process your request", http.StatusInternalServerError)
	}
}
