package api

import (
	"encoding/json"
	"net/http"

	"github.com/paddyoneill/go-svelte/internal/store"
)

type Server struct {
	serverStore store.Store
}

func NewServer(s store.Store) Server {
	return Server{
		serverStore: s,
	}
}

func (s Server) GetPing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Pong{
		Result: "pong",
	})
}

func (s Server) PostEntry(w http.ResponseWriter, r *http.Request) {
	var entry Entry

	err := json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, ok := s.serverStore[entry.Id]; ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	s.serverStore[entry.Id] = entry.Entry
	w.WriteHeader(http.StatusOK)
}

func (s Server) GetEntry(w http.ResponseWriter, r *http.Request) {
	var entries EntryList
	for id, entry := range s.serverStore {
		entries = append(entries, Entry{Id: id, Entry: entry})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(entries)
}
