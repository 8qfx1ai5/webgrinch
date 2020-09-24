package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func init() {
	//http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func main() {
	url, err := url.Parse("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(url)
	http.ListenAndServe(":80", proxy)
}
