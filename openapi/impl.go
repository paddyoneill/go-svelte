package api

import (
	"encoding/json"
	"net/http"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (s Server) GetPing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Pong{
		Result: "pong",
	})
}
