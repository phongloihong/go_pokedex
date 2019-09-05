package main

import (
	"context"
	"fmt"

	PokemonModel "github.com/phongloihong/go_tag/models"
	mongoService "github.com/phongloihong/go_tag/services/db"
	pokemonApi "github.com/phongloihong/go_tag/services/pokeApi"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	connection := mongoService.Connect()
	defer connection.Client.Disconnect(context.Background())

	getPokemonFromApi(*connection.Database)
}

func getPokemonFromApi(db mongo.Database) {
	pokemonAPI := pokemonApi.PokemonAPI{}
	pokemonAPI.GetPokemons()

	pokemonModel := PokemonModel.Create(db)
	for _, v := range pokemonAPI.PokemonList {
		isExisted, _ := pokemonModel.GetByName(v.Name)
		if isExisted {
			fmt.Printf("Pokemon %s existed\n", v.Name)
			continue
		}

		fmt.Printf("Insert pokemon %s\n", v.Name)
		pokemonModel.InsertPokemon(v)
	}
}
