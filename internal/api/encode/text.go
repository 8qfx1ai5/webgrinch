package encode

import (
	"net/http"

	"github.com/8qfx1ai5/viewcrypt/configs"
	"github.com/8qfx1ai5/viewcrypt/internal/api"
	"github.com/8qfx1ai5/viewcrypt/internal/encodetext"
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
		api.Error(w, "request params invalid", http.StatusNotAcceptable, err)
		return
	}

	content := r.Form.Get("content")
	keyFrom := r.Form.Get("from")
	keyTo := r.Form.Get("to")

	if keyFrom == "" {
		keyFrom = configs.APIDefaultKeyFrom
	}
	if keyTo == "" {
		keyTo = configs.APIDefaultKeyTo
	}

	// run encoding
	encoded, err := encodetext.Run(content, keyFrom, keyTo)
	if err != nil {
		api.Error(w, "encoding failed", http.StatusInternalServerError, err)
		return
	}

	// create response
	api.Success(w, api.Response{Payload: encoded})
}
