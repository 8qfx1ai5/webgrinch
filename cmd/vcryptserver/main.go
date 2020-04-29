package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/8qfx1ai5/viewcrypt/internal/apiencode"
	"github.com/8qfx1ai5/viewcrypt/internal/swagger"
)

// configuration for the api and encoder
const (
	baseURL string = ""
)

// initialize webserver and route to the controllers
func main() {

	var cliArguments = handleCliArguments()

	// http.Handle("/swagger/", http.StripPrefix("/swagger/", swagger.FileServer()))
	http.Handle("/swagger/", http.StripPrefix("/swagger/", swagger.FileServer()))
	http.HandleFunc(baseURL+"/encode", apiencode.RouteHandler)

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
