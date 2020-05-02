package key

import (
	"fmt"
	"net/http"

	"github.com/8qfx1ai5/viewcrypt/internal/api"
	"github.com/8qfx1ai5/viewcrypt/internal/types/enkey"
)

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
		api.Success(w, api.Response{Payload: fmt.Sprint(newKey)})
		return
	default:
		api.Error(w, "try an other http method or have a look into our api documentation", http.StatusMethodNotAllowed, nil)
		return
	}
}
