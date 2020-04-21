package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/8qfx1ai5/viewcrypt/internal/encode"
)

// configuration for the api and encoder
const (
	baseURL string = ""
	keyFrom string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	keyTo   string = "aFMkZVwKEWsjUQdgYfuIpNGSDnyxPehiLTRbCoqvXmAzBcrltHJO"
)

// initialize webserver and route to the controllers
func main() {
	var cliArguments = handleCliArguments()
	http.HandleFunc(baseURL+"/encode", encodeController)
	http.ListenAndServe(fmt.Sprintf(":%s", cliArguments.apiPort), nil)
}

// Response the structure of the json response used by the api
type Response struct {
	Encoded string
}

func encodeController(w http.ResponseWriter, r *http.Request) {
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

// all expected cli arguments as return type
type cliArguments struct {
	apiPort string
}

// parse the cli arguments
func handleCliArguments() (out cliArguments) {
	apiPort := flag.String("p", "8888", "use this port for the web server")
	flag.Parse()
	out.apiPort = *apiPort

	return
}
