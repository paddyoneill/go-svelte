package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	api "github.com/paddyoneill/go-svelte/openapi"
)

func main() {
	server := api.NewServer()
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		viteURL, _ := url.Parse("http://localhost:5173")
		reverseViteProxy := httputil.NewSingleHostReverseProxy(viteURL)
		reverseViteProxy.ServeHTTP(w, r)
	})
	handler := api.HandlerFromMux(server, router)

	if err := http.ListenAndServe("localhost:8080", handler); err != nil {
		panic(err)
	}
}
