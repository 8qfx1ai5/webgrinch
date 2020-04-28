package api

import (
	"fmt"
	"log"
	"net/http"
)

// Error function modifies the http.Error function
func Error(w http.ResponseWriter, message string, code int, err error) {
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Server", "The Viewcrypt Go Sebserver")
	w.WriteHeader(code)
	fmt.Fprintln(w, fmt.Sprintf("{\"message\":\"%s\"}", message))
}
