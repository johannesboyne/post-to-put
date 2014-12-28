package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var in = flag.String("in", "8080", "default incoming port")
var out = flag.String("out", "http://127.0.0.1:8500", "default outgoing URI")

func main() {
	flag.Parse()
	remote, err := url.Parse(*out)
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.HandleFunc("/", handler(proxy))
	err = http.ListenAndServe(":"+*in, nil)
	if err != nil {
		panic(err)
	}
}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := http.NewRequest("PUT", *out+r.URL.String(), r.Body)
		if err != nil {
			log.Fatal(err)
		}
		p.ServeHTTP(w, req)
	}
}
