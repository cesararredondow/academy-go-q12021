package handlers

import (
	"net/http"

	"github.com/cesararredondow/course/second_deliverable/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

type UseCase interface {
	GetPokemons(string) ([]*models.Pokemon, error)
	GetPokemon(string) (*models.PokemonResponse, error)
}

// Pokemons struct
type Pokemons struct {
	useCase UseCase
	logger  *logrus.Logger
	render  *render.Render
}

// New returns a controller
func New(
	u UseCase,
	logger *logrus.Logger,
	r *render.Render,
) *Pokemons {
	return &Pokemons{u, logger, r}
}

// GetPokemonsFromApi is the handler to get all the pokemons by quantity
func (p *Pokemons) GetPokemonsFromApi(w http.ResponseWriter, r *http.Request) {
	logger := p.logger.WithFields(logrus.Fields{
		"func": "Get Pokemons",
	})
	logger.Debug("in")
	quantity := r.FormValue("quantity")
	if quantity == "" {
		quantity = "1"
	}
	pokemons, err := p.useCase.GetPokemons(quantity)
	if err != nil {
		p.render.JSON(w, http.StatusInternalServerError, pokemons)
	}

	p.render.JSON(w, http.StatusOK, pokemons)

}

// GetPokemonsFromApi is the handler to get the pokemon by id
func (p *Pokemons) GetPokemonFromApi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pokemonID := vars["id"]
	logger := p.logger.WithFields(logrus.Fields{
		"func": "Get Pokemon",
		"id":   pokemonID,
	})
	logger.Debug("in")

	pokemon, err := p.useCase.GetPokemon(pokemonID)

	if err == nil {
		p.render.JSON(w, http.StatusOK, pokemon)
	}
	p.render.JSON(w, http.StatusNotFound, pokemon)
}
