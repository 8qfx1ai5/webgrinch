package encode

import (
	"fmt"
	"net/http"

	"github.com/8qfx1ai5/viewcrypt/configs"
	"github.com/8qfx1ai5/viewcrypt/internal/api"
	"github.com/8qfx1ai5/viewcrypt/internal/encodehtml"
)

// HTMLHandler handles the encoding feature
func HTMLHandler(w http.ResponseWriter, r *http.Request) {

	// check html method
	if r.Method != "POST" {
		api.Error(w, "webserver doesn't support this http method", http.StatusMethodNotAllowed, nil)
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

	css := r.Form.Get("css")
	keyFrom := r.Form.Get("from")
	keyTo := r.Form.Get("to")

	if keyFrom == "" {
		keyFrom = configs.APIDefaultKeyFrom
	}
	if keyTo == "" {
		keyTo = configs.APIDefaultKeyTo
	}

	// run encoding
	encoded, err := encodehtml.Run(payload, keyFrom, keyTo, css)
	if err != nil {
		api.Error(w, "encoding failed", http.StatusInternalServerError, fmt.Errorf(fmt.Sprintf("payload='%s' keyFrom='%s' keyTo='%s' css='%s'\n", payload, keyFrom, keyTo, css), err))
		return
	}

	// create response
	api.Success(w, api.Response{Payload: encoded})
}
