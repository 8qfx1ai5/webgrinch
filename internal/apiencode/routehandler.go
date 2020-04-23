package apiencode

import (
	"encoding/json"
	"net/http"

	"github.com/8qfx1ai5/viewcrypt/internal/encode"
)

const (
	keyFromDefault string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	keyToDefault   string = "BCDEFGHIJKLMNOPQRSTUVWXYZAbcdefghijklmnopqrstuvwxyza"
)

// Response the structure of the json response used by the api
type Response struct {
	Encoded string
}

// RouteHandler handles the encoding feature
func RouteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "The Viewcrypt Go Sebserver")

	if r.Method != "POST" {
		// TODO: logging?
		http.Error(w, "webserver doesn't support this method", http.StatusMethodNotAllowed)
		return
	}

	// get post params
	err := r.ParseForm()
	if err != nil {
		// TODO: logging?
		http.Error(w, "request params invalid", http.StatusNotAcceptable)
		return
	}

	content := r.Form.Get("content")
	css := r.Form.Get("css")
	keyFrom := r.Form.Get("from")
	keyTo := r.Form.Get("to")

	if keyFrom == "" {
		keyFrom = keyFromDefault
	}
	if keyTo == "" {
		keyTo = keyToDefault
	}

	// run encoding
	encoded, err := encode.HTML(content, keyFrom, keyTo, css)
	if err != nil {
		// TODO: logging?
		http.Error(w, "encoding failed", http.StatusInternalServerError)
		return
	}

	// create response
	response := Response{encoded}

	// convert into json
	js, err := json.Marshal(response)
	if err != nil {
		// TODO: logging?
		http.Error(w, "response conversion failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write(js)
}
