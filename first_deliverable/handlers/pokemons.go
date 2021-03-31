package handlers

import (
	"net/http"

	"github.com/cesararredondow/course/first_deliverable/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

type UseCase interface {
	GetPokemons() ([]*models.Pokemon, error)
	GetPokemon(string) (*models.Pokemon, error)
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

//GetPokemons this is the endpoint to get all the pokemons from the csv
func (p *Pokemons) GetPokemons(w http.ResponseWriter, r *http.Request) {
	logger := p.logger.WithFields(logrus.Fields{
		"func": "Get Pokemons",
	})
	logger.Debug("in")
	pokemons, err := p.useCase.GetPokemons()
	if err != nil {
		p.logger.Error(err)
		p.render.JSON(w, http.StatusInternalServerError, pokemons)
	}

	p.render.JSON(w, http.StatusOK, pokemons)

}

//GetPokemon this is the endpoint and specific pokemon from the csv, it's selected by their id
func (p *Pokemons) GetPokemon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pokemonID := vars["id"]
	logger := p.logger.WithFields(logrus.Fields{
		"func": "Get Pokemon",
		"id":   pokemonID,
	})
	logger.Debug("in")

	pokemon, err := p.useCase.GetPokemon(pokemonID)

	if err == nil {
		p.logger.Error(err)
		p.render.JSON(w, http.StatusOK, pokemon)
	}
	p.render.JSON(w, http.StatusNotFound, pokemon)
}
