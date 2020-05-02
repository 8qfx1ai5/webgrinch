package key

import (
	"net/http"

	"github.com/8qfx1ai5/viewcrypt/configs"
	"github.com/8qfx1ai5/viewcrypt/internal/api"
	"github.com/8qfx1ai5/viewcrypt/internal/types/enkey"
)

// Response the structure of the json response used by the api
type Response struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// Handler handles the features to the encoding keys
func Handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		// handle post params
		err := r.ParseForm()
		if err != nil {
			api.Error(w, "request params invalid", http.StatusNotAcceptable, err)
			return
		}
		regex := r.Form.Get("regex")
		if regex == "" {
			regex = configs.APIDefaultKeyRegex
		}
		var newKey = enkey.Key{}
		ok, err := newKey.UseRegex(regex)
		if err != nil {
			api.Error(w, "something went wrong", http.StatusInternalServerError, err)
			return
		}
		if !ok {
			api.Error(w, "regex not supported", http.StatusBadRequest, nil)
			return
		}
		api.Success(w, Response{From: newKey.GetFrom(), To: newKey.GetTo()})
		return
	default:
		api.Error(w, "try an other http method or have a look into our api documentation", http.StatusMethodNotAllowed, nil)
		return
	}
}
