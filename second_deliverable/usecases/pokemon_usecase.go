package usecases

import "github.com/cesararredondow/course/second_deliverable/models"

type UseCase struct {
	service Service
}

type Service interface {
	GetPokemons(string) ([]*models.Pokemon, error)
	GetPokemon(string) (*models.PokemonResponse, error)
}

func New(service Service) *UseCase {
	return &UseCase{service}
}

//GetPokemons is the usecase to get the information
func (u *UseCase) GetPokemons(quantity string) ([]*models.Pokemon, error) {
	resp, err := u.service.GetPokemons(quantity)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

////GetPokemon is the usecase to get the information
func (u *UseCase) GetPokemon(pokemonID string) (*models.PokemonResponse, error) {
	resp, err := u.service.GetPokemon(pokemonID)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
