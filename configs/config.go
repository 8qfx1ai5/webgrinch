package configs

import "net/http"

const (
	// APIDefaultKeyFrom used in the encode apis, if no key is specified
	APIDefaultKeyFrom string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	// APIDefaultKeyTo used in the encode apis, if no key is specified
	APIDefaultKeyTo string = "BCDEFGHIJKLMNOPQRSTUVWXYZAbcdefghijklmnopqrstuvwxyza"

	// APIServerHeaderName is shown in the response header as "server: "
	APIServerHeaderName string = "The Viewcrypt Golang Server"

	// APIDefaultKeyRegex used for key generation if nothing specified
	APIDefaultKeyRegex string = "[A-Za-z0-9]"
)

// ServerSetDefaultHeaders must be called in every api handler
func ServerSetDefaultHeaders(w http.ResponseWriter) {
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Server", APIServerHeaderName)
}
