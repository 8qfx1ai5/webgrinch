package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"path/filepath"

	"github.com/8qfx1ai5/webgrinch/internal/api/encode"
	"github.com/8qfx1ai5/webgrinch/internal/api/key"
	"github.com/8qfx1ai5/webgrinch/internal/swagger"
)

// configuration for the api and encoder
const (
	baseURL string = ""
)

var tlsCertPath = filepath.FromSlash("/certs/cert.pem")
var tlsCertKeyPath = filepath.FromSlash("/certs/privkey.pem")

// initialize webserver and route to the controllers
func main() {

	var cliArguments = handleCliArguments()

	go redirectPort80(cliArguments.apiPort)

	http.Handle("/", http.FileServer(http.Dir(filepath.Join("webgrinch", "web", "static", "example"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join("webgrinch", "web", "static")))))
	http.Handle("/api/", http.StripPrefix("/api/", swagger.FileServer()))
	http.HandleFunc(baseURL+"/api/encode/html", encode.HTMLHandler)
	http.HandleFunc(baseURL+"/api/encode/text", encode.TextHandler)
	http.HandleFunc(baseURL+"/api/key", key.Handler)

	if handleTLSCert() {
		log.Printf("Start server in TSL mode.")
		err := http.ListenAndServeTLS(fmt.Sprintf(":%s", cliArguments.apiPort), tlsCertPath, tlsCertKeyPath, nil)
		if err != nil {
			log.Print(err)
		}
	} else {
		log.Printf("TSL cert missing!!! Start without TSL. (cert: \"%s\" - key: \"%s\")", tlsCertPath, tlsCertKeyPath)
		err := http.ListenAndServe(fmt.Sprintf(":%s", cliArguments.apiPort), nil)
		if err != nil {
			log.Print(err)
		}
	}
}

func handleTLSCert() bool {
	// check that files exist
	if dat, err := ioutil.ReadFile(tlsCertPath); err != nil || len(dat) == 0 {
		log.Print(err)
		return false
	}
	if dat, err := ioutil.ReadFile(tlsCertKeyPath); err != nil || len(dat) == 0 {
		log.Print(err)
		return false
	}
	return true
}

// all expected cli arguments as return type
type cliArguments struct {
	apiPort string
}

// parse the cli arguments
func handleCliArguments() (out cliArguments) {
	apiPort := flag.String("p", "8888", "use this port for the web server")
	flag.Parse()
	out.apiPort = *apiPort

	return out
}

func redirectPort80(tlsPort string) {
	log.Println("Start redirect server on port 80.")
	httpSrv := http.Server{
		Addr: ":80",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			host, _, err := net.SplitHostPort(r.Host)
			if err != nil {
				host = r.Host
			}
			u := r.URL
			u.Host = net.JoinHostPort(host, tlsPort)
			u.Scheme = "https"
			log.Printf("u.String()=%s", u.String())
			http.Redirect(w, r, u.String(), http.StatusMovedPermanently)
		}),
	}
	log.Println(httpSrv.ListenAndServe())
}
