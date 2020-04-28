package apiencode

import (
	"net/http"

	"github.com/8qfx1ai5/viewcrypt/internal/api"
	"github.com/8qfx1ai5/viewcrypt/internal/encode"
)

const (
	keyFromDefault string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	keyToDefault   string = "BCDEFGHIJKLMNOPQRSTUVWXYZAbcdefghijklmnopqrstuvwxyza"
)

// RouteHandler handles the encoding feature
func RouteHandler(w http.ResponseWriter, r *http.Request) {

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
		api.Error(w, "encoding failed", http.StatusInternalServerError, err)
		return
	}

	// create response
	api.Success(w, api.Response{Content: encoded})
}
