package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	target, err := url.Parse("localhost:14265")
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	http.HandleFunc("/", handler(proxy))
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func handler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(writter http.ResponseWriter, req *http.Request) {
		writter.Header().Set("Content-Type", "application/json")
		writter.Header().Set("X-IOTA-API-Version", "1")

		// Log
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}

		req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		log.Println(req.Body)

		proxy.ServeHTTP(writter, req)
	}
}
