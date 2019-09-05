package pokemonApi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	PokemonModel "github.com/phongloihong/go_tag/models"
)

/**
 * Struct for pokeAPi response
 */
type PokemonAPI struct {
	PokemonList []PokemonModel.Pokemon `json:"results"`
}

func (pApi *PokemonAPI) GetPokemons() *PokemonAPI {
	response, err := http.Get("https://pokeapi.co/api/v2/pokemon?offset=20&limit=20")
	if err != nil {
		panic(err)
	}

	data, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(data, &pApi)

	return pApi
}
