package services

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/cesararredondow/course/second_deliverable/models"
)

type Service struct {
	writer     *csv.Writer
	pokemonAPI string
	pathFile   string
}

func New(writer *csv.Writer, pokemonAPI string, pathFile string) (*Service, error) {
	return &Service{writer, pokemonAPI, pathFile}, nil
}

//GetPokemons get an amount of pokemons that you request from pokemon API
func (s *Service) GetPokemons(quantity string) ([]*models.Pokemon, error) {
	pokemons := []*models.Pokemon{}
	url := "/pokemon/?limit="
	url += quantity

	response, err := http.Get(s.pokemonAPI + url)
	if err != nil {
		return nil, errors.New("Error")
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("Error")
	}

	jsonData := models.Response{}
	json.Unmarshal(responseData, &jsonData)

	for i, pokemon := range jsonData.Results {
		p := new(models.Pokemon)
		p.ID = i + 1
		p.Name = pokemon.Name
		p.URL = pokemon.URL
		pokemons = append(pokemons, p)
	}

	if errCSV := s.writeJSONInCSV(jsonData); errCSV != nil {
		return nil, errCSV
	}

	return pokemons, nil
}

//GetPokemon get the pokemons that you request from pokemon API
func (s *Service) GetPokemon(pokemonID string) (*models.PokemonResponse, error) {
	response, err := http.Get(s.pokemonAPI + "/pokemon/" + pokemonID)

	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("Error")
	}

	jsonData := models.PokemonResponse{}
	json.Unmarshal(responseData, &jsonData)

	return &jsonData, nil
}

func (s *Service) writeJSONInCSV(jsonData models.Response) error {
	_, err := os.Stat(s.pathFile)
	if os.IsNotExist(err) {
		s.writer.Write([]string{"ID", "NAME"})
	}
	for i, p := range jsonData.Results {
		var row []string
		row = append(row, strconv.Itoa(i+1))
		row = append(row, p.Name)
		s.writer.Write(row)
	}

	s.writer.Flush()
	return nil
}
