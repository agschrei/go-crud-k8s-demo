package main

import (
	"encoding/json"
	"net/http"
)

type AppStatus struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

func (app *application) statusHandler(w http.ResponseWriter, r *http.Request) {
	currentStatus := AppStatus{
		Status:  "UP",
		Version: version,
	}

	payload, err := json.MarshalIndent(currentStatus, "", "\t")
	if err != nil {
		app.config.Logger.Println(err)
		app.errorJSON(w, err, 500)
		return
	}

	app.config.Logger.Print("GET " + r.RequestURI + " 200 OK")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}

func (app *application) getAllAirports(w http.ResponseWriter, r *http.Request) {
	repo := app.repository
	airports, err := repo.GetAllAirports()

	if err != nil {
		app.config.Logger.Printf("Encountered error trying to get from DB: %s", err)
		app.errorJSON(w, err, 500)
		return
	}

	app.config.Logger.Print("GET " + r.RequestURI + " 200 OK")

	app.writeJSON(w, 200, airports)
}

func (app *application) getAllAircraft(w http.ResponseWriter, r *http.Request) {
	repo := app.repository
	aircraft, err := repo.GetAllAircraft()

	if err != nil {
		app.config.Logger.Printf("Encountered error trying to get from DB: %s", err)
		app.errorJSON(w, err, 500)
		return
	}

	app.config.Logger.Print("GET " + r.RequestURI + " 200 OK")

	app.writeJSON(w, 200, aircraft)
}
