package handlers

import (
	"net/http"

	"github.com/cesararredondow/course/third_deriverable/models"

	"github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

type UseCase interface {
	GetPokemonsConcurrency(odd string, quantity string, workersNumer string) ([]*models.Pokemon, error)
}

// DoSomenthing struct
type DoSomenthing struct {
	useCase UseCase
	logger  *logrus.Logger
	render  *render.Render
}

// New returns a controller
func New(
	u UseCase,
	logger *logrus.Logger,
	r *render.Render,
) *DoSomenthing {
	return &DoSomenthing{u, logger, r}
}

//GetPokemonsConcurrecny is a handler to get the pokemons
func (d *DoSomenthing) GetPokemonsConcurrecny(w http.ResponseWriter, r *http.Request) {
	odd := r.FormValue("odd")
	quantity := r.FormValue("quantity")
	nWorkers := r.FormValue("numberWorkers")

	if odd == "" || quantity == "" || nWorkers == "" {
		d.logger.Error("Missing query params")
		d.render.JSON(w, http.StatusBadRequest, "Missing query params")
		return
	}

	tmp, err := d.useCase.GetPokemonsConcurrency(odd, quantity, nWorkers)
	if err != nil {
		d.logger.Error(err)
		d.render.JSON(w, http.StatusInternalServerError, tmp)
	}
	d.render.JSON(w, http.StatusOK, tmp)
}
