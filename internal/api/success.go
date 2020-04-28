package api

import (
	"encoding/json"
	"net/http"
)

// Response the structure of the json response used by the api
type Response struct {
	Content string
}

// Success function writes response
func Success(w http.ResponseWriter, response Response) {
	// convert into json
	js, err := json.Marshal(response)
	if err != nil {
		Error(w, "response conversion failed", http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Server", "The Viewcrypt Go Sebserver")
	w.WriteHeader(200)
	w.Write(js)
}
