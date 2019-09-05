package PokemonModel

import (
	mongoService "github.com/phongloihong/go_tag/services/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Pokemon struct {
	Name string `json:"name" bson:"name"`
	Url  string `json:"url" bson:"url"`
}

type PokemonCollection struct {
	Collection *mongo.Collection
}

func Create(db mongo.Database) *PokemonCollection {
	return &PokemonCollection{db.Collection("Pokemons")}
}

func (p *PokemonCollection) InsertPokemon(v Pokemon) {
	ctx, _ := mongoService.ContextTimeOut()
	_, err := p.Collection.InsertOne(ctx, v)
	if err != nil {
		panic(err)
	}
}

func (p *PokemonCollection) GetByName(pokemonName string) (bool, Pokemon) {
	var pokemon Pokemon
	ctx, _ := mongoService.ContextTimeOut()

	err := p.Collection.FindOne(ctx, bson.M{"name": pokemonName}).Decode(&pokemon)
	if err != nil {
		panic(err)
	}

	return pokemon != (Pokemon{}), pokemon
}
