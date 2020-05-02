package api

import (
	"encoding/json"
	"net/http"

	"github.com/8qfx1ai5/viewcrypt/configs"
)

// Response the structure of the json response used by the api
type Response struct {
	Payload string `json:"payload"`
}

// Success function writes response
func Success(w http.ResponseWriter, response interface{}) {
	// convert into json
	js, err := json.Marshal(response)
	if err != nil {
		Error(w, "response conversion failed", http.StatusInternalServerError, err)
		return
	}

	configs.ServerSetDefaultHeaders(w)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(200)
	w.Write(js)
}

// TODO: use interface for payload (more generic) or find route based solution
