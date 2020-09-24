package encode

import (
	"net/http"

	"github.com/8qfx1ai5/webgrinch/configs"
	"github.com/8qfx1ai5/webgrinch/internal/api"
	"github.com/8qfx1ai5/webgrinch/internal/encodetext"
)

// TextHandler handles the encoding feature
func TextHandler(w http.ResponseWriter, r *http.Request) {

	// check html method
	if r.Method != "POST" {
		api.Error(w, "webserver doesn't support this method", http.StatusMethodNotAllowed, nil)
		return
	}

	// handle post params
	err := r.ParseForm()
	if err != nil {
		api.Error(w, "params could not be parsed", http.StatusInternalServerError, err)
		return
	}

	payload := r.Form.Get("payload")
	if payload == "" {
		api.ParamError(w, "payload", "payload is empty", nil)
		return
	}

	keyFrom := r.Form.Get("from")
	keyTo := r.Form.Get("to")

	if keyFrom == "" {
		keyFrom = configs.APIDefaultKeyFrom
	}
	if keyTo == "" {
		keyTo = configs.APIDefaultKeyTo
	}

	// run encoding
	encoded, err := encodetext.Run(payload, keyFrom, keyTo)
	if err != nil {
		api.Error(w, "encoding failed", http.StatusInternalServerError, err)
		return
	}

	// create response
	api.Success(w, api.Response{Payload: encoded})
}
