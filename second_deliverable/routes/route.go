package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Controller interface {
	GetPokemonsFromApi(w http.ResponseWriter, r *http.Request)
	GetPokemonFromApi(w http.ResponseWriter, r *http.Request)
}

func New(r *mux.Router, controller Controller) {
	api := r.PathPrefix("/api/v2").Subrouter()
	api.HandleFunc("/pokemons", controller.GetPokemonsFromApi).Methods(http.MethodGet).Name("get_Pokemons")
	api.HandleFunc("/pokemon/{pokemonNumber}", controller.GetPokemonFromApi).Methods(http.MethodGet).Name("get_pokemon")
}
