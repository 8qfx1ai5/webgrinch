package swagger

import (
	"fmt"
	"log"
	"net/http"

	"github.com/8qfx1ai5/webgrinch/configs"

	"github.com/rakyll/statik/fs"
	// import compiled swagger ui files
	_ "github.com/8qfx1ai5/webgrinch/third_party/swagger-ui/statik"
)

// FileServer create new file server for swagger ui files
func FileServer() http.Handler {
	statikFS, err := fs.New()
	if err != nil {
		log.Println(err)
		return errHandler
	}

	return http.FileServer(statikFS)
}

type errorHandler struct{}

var errHandler = errorHandler{}

const (
	errorFileName    string = "index.html"
	errorFileContent string = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Server error, Swagger-UI not found</title>
</head>
<body>
	<h1>Upps... Swagger-UI not found.</h1>
</body>
</html>
`
)

func (errorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	configs.ServerSetDefaultHeaders(w)
	w.Header().Set("Content-Type", "text/html")

	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintln(w, errorFileContent)
}
