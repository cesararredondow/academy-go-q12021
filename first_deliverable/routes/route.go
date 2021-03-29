package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Controller interface {
	GetPokemons(w http.ResponseWriter, r *http.Request)
	GetPokemon(w http.ResponseWriter, r *http.Request)
}

func New(c Controller, r *mux.Router) {
	first := r.PathPrefix("/api/v1").Subrouter()
	first.HandleFunc("/pokemons", c.GetPokemons).Methods(http.MethodGet).Name("get_Pokemons")
	first.HandleFunc("/pokemon/{id}", c.GetPokemon).Methods(http.MethodGet).Name("get_Pokemon")
}
