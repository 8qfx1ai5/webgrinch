package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/8qfx1ai5/webgrinch/internal/api/encode"
	"github.com/8qfx1ai5/webgrinch/internal/api/key"
	"github.com/8qfx1ai5/webgrinch/internal/swagger"
)

// configuration for the api and encoder
const (
	baseURL string = ""
)

var tlsCertPath = filepath.FromSlash("/webgrinch/tmp/certs/cert.pem")
var tlsCertKeyPath = filepath.FromSlash("/webgrinch/tmp/certs/privkey.pem")

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
	// get the TLS cert credentials from the environment vars
	var tlsCert = os.Getenv("TLSCERT")
	if tlsCert == "" {
		log.Println("environment variable TLSCERT is empty")
	}
	var tlsCertKey = os.Getenv("TLSCERTKEY")
	if tlsCertKey == "" {
		log.Println("environment variable TLSCERTKEY is empty")
	}
	// create temporary directory for the cert files
	if _, err := os.Stat(filepath.Dir(tlsCertPath)); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(tlsCertPath), 0700)
		if err != nil {
			log.Println(err)
		}
	}
	// create the cert files with the input from the environment vars
	tlsCertFile, _ := os.Create(tlsCertPath)
	_, _ = tlsCertFile.WriteString(strings.Replace(tlsCert, "#", "\n", -1))
	if err := tlsCertFile.Sync(); err != nil {
		log.Print(err)
	}
	tlsCertKeyFile, _ := os.Create(tlsCertKeyPath)
	_, _ = tlsCertKeyFile.WriteString(strings.Replace(tlsCertKey, "#", "\n", -1))
	if err := tlsCertKeyFile.Sync(); err != nil {
		log.Print(err)
	}

	// check that files exist
	if _, err := os.Stat(tlsCertPath); err != nil {
		log.Print(err)
		return false
	}
	if _, err := os.Stat(tlsCertKeyPath); err != nil {
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
