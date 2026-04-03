package main

import (
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type PingResponse struct {
	Result string `json:"result"`
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		viteURL, _ := url.Parse("http://localhost:5173")
		reverseViteProxy := httputil.NewSingleHostReverseProxy(viteURL)
		reverseViteProxy.ServeHTTP(w, r)
	})
	router.HandleFunc("GET /ping", handleGetPing)

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}

func handleGetPing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PingResponse{Result: "pong"})
}
