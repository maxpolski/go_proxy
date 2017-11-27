package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	backend := ":9000"
	proxy := ":9001"

	u, _ := url.Parse("http://localhost" + backend)
	http.Handle("/", httputil.NewSingleHostReverseProxy(u))

	log.Fatal(http.ListenAndServe(proxy, nil))
}
