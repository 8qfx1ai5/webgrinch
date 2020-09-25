package main

import (
	"flag"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/8qfx1ai5/webgrinch/internal/api/encode"
	"github.com/8qfx1ai5/webgrinch/internal/api/key"
	"github.com/8qfx1ai5/webgrinch/internal/swagger"
)

// configuration for the api and encoder
const (
	baseURL string = ""
)

// initialize webserver and route to the controllers
func main() {

	var cliArguments = handleCliArguments()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join("webgrinch", "static")))))
	http.Handle("/", http.StripPrefix("/", swagger.FileServer()))
	http.HandleFunc(baseURL+"/api/encode/html", encode.HTMLHandler)
	http.HandleFunc(baseURL+"/api/encode/text", encode.TextHandler)
	http.HandleFunc(baseURL+"/api/key", key.Handler)

	http.ListenAndServe(fmt.Sprintf(":%s", cliArguments.apiPort), nil)
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

	return out
}
