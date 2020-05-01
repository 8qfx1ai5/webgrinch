package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/8qfx1ai5/viewcrypt/configs"
)

// Error function modifies the http.Error function
func Error(w http.ResponseWriter, message string, code int, err error) {
	configs.ServerSetDefaultHeaders(w)

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprintln(w, fmt.Sprintf("{\"message\":\"%s\"}", message))
}

// TODO: handle Server Errors
// TODO: handle Special Client errors like 404, 406
