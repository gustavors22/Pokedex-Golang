package pokedex

import (
	"errors"
	"fmt"
	"strconv"
)

type pokemon struct {
	Name 		string	
	PokemonType string	
	Id			int		
}

var pokedex []pokemon 

func CreatePokemon(name string, pokemonType string, id int) {
	pokemon := pokemon{name, pokemonType, id}
	pokedex = append(pokedex, pokemon)
}

func GetPokedex(){
	if len(pokedex) == 0{
		fmt.Printf("Pokedex is empty. Add pokemons!!\n")

	} else {
		for _, pkm := range(pokedex){
			fmt.Printf("\n\n\nName: %s\nType: %s\nId: %d", pkm.Name, pkm.PokemonType, pkm.Id)
		}
	}
}

func GetPokemon(id int) error{
	if len(pokedex) == 0{
		return errors.New("pokedex is empty. Add pokemons")
	}

	for i := range(pokedex){
		if pokedex[i].Id == id {
			fmt.Printf("\nName: %s\nType: %s\nId: %d\n", pokedex[i].Name, pokedex[i].PokemonType, pokedex[i].Id)
			break
		}

		if i == len(pokedex) - 1 && pokedex[i].Id != id {
			return errors.New("pokemon not registed")
		}
	}

	return nil
}

func DeletePokemon(id int) error {
	if len(pokedex) == 0{
		return errors.New("pokedex is empty. Add pokemons")
	}

	for i := range(pokedex){
		if pokedex[i].Id == id {
			pokedex = pokedex[:i+copy(pokedex[i:], pokedex[i+1:])]
			break
		}

		if i == len(pokedex) - 1 && pokedex[i].Id != id {
			return errors.New("pokemon not registed")
		}
	}

	return nil
}

func UpdatePokemon(id int, pokemonData []string) error {
	if len(pokedex) == 0{
		return errors.New("pokedex is empty. Add pokemons")
	}

	for i := range(pokedex){
		if pokedex[i].Id == id {
			pokedex[i].Name = pokemonData[0]
			pokedex[i].PokemonType = pokemonData[1]
			pokedex[i].Id , _= strconv.Atoi(pokemonData[2])
			break
		}

		if i == len(pokedex) - 1 && pokedex[i].Id != id {
			return errors.New("pokemon not registed")
		}
	}

	return nil
}