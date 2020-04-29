package swagger

import (
	"fmt"
	"log"
	"net/http"

	// import compiled swagger ui files
	_ "github.com/8qfx1ai5/viewcrypt/third_party/swagger-ui/statik"
	"github.com/rakyll/statik/fs"
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
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Server", "The Viewcrypt Go Sebserver")
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintln(w, errorFileContent)
}
