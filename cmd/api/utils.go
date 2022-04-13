package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}) error {
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(payload)

	return nil
}

func (app *application) errorJSON(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	type ErrorPayload struct {
		ErrorMsg string `json:"error_msg"`
	}

	errPayload := ErrorPayload{
		ErrorMsg: err.Error(),
	}

	app.writeJSON(w, statusCode, errPayload)
}
