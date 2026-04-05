package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/paddyoneill/go-svelte/internal/store"
	api "github.com/paddyoneill/go-svelte/openapi"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		viteURL, _ := url.Parse("http://localhost:5173")
		reverseViteProxy := httputil.NewSingleHostReverseProxy(viteURL)
		reverseViteProxy.ServeHTTP(w, r)
	})

	server := api.NewServer(store.Store{})
	handler := api.HandlerWithOptions(server, api.StdHTTPServerOptions{
		BaseRouter: router,
		BaseURL:    "/api",
	})

	if err := http.ListenAndServe("localhost:8080", handler); err != nil {
		panic(err)
	}
}
