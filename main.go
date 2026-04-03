package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		viteURL, _ := url.Parse("http://localhost:5173")
		reverseViteProxy := httputil.NewSingleHostReverseProxy(viteURL)
		reverseViteProxy.ServeHTTP(w, r)
	})

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}
