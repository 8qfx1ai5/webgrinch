package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/8qfx1ai5/webgrinch/configs"
)

// Error function modifies the http.Error function
func Error(w http.ResponseWriter, hint string, code int, err error) {
	configs.ServerSetDefaultHeaders(w)

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	// don't use json formatting to omit additional errors
	switch code {
	case 500:
		fmt.Fprintln(w, fmt.Sprintf("{\"hint\":\"%s\"}", configs.APIDefaultServerError500Hint))
	default:
		fmt.Fprintln(w, fmt.Sprintf("{\"hint\":\"%s\"}", hint))
	}

}

// ParamError function modifies the http.Error function
func ParamError(w http.ResponseWriter, pname string, hint string, err error) {
	configs.ServerSetDefaultHeaders(w)

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	// don't use json formatting to omit additional errors
	fmt.Fprintln(w, fmt.Sprintf("{\"pname\":\"%s\",\"hint\":\"%s\"}", pname, hint))
}
