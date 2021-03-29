package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Controller interface {
	GetPokemonsConcurrecny(w http.ResponseWriter, r *http.Request)
}

func New(c Controller, r *mux.Router) {
	final := r.PathPrefix("/api/final").Subrouter()
	final.HandleFunc("/", c.GetPokemonsConcurrecny).Methods(http.MethodGet).Name("get_dosomenthing")
}
