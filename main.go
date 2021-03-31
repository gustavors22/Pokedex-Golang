package main

import (
	"fmt"
	"project/pokedex"
) 


func main(){
	
	banner()

	for {
		menu()

		var option int
		fmt.Scanf("%d", &option)

		if option == 1 {
			var name string
			var typePokemon string
			var id int
			
			fmt.Println("Write pokemon name, pokemon type and pokemon id. Before press ENTER")
			fmt.Scanf("%s %s %d", &name, &typePokemon, &id)

			pokedex.CreatePokemon(name,typePokemon, id)

		}
		
		if option == 2 {
			var id int

			fmt.Println("Write a pokemon id for search")
			fmt.Scanf("%d", &id)

			err := pokedex.GetPokemon(id)

			if err != nil {
				fmt.Println(err.Error())
			}
		}

		if option == 3 {
			pokedex.GetPokedex()
			fmt.Println("")
		}

		if option == 4 {
			var pokemonData []string
			var id int

			fmt.Println("Write a pokemon id for update")
			fmt.Scanf("%d", &id)

			fmt.Println("Write pokemon name, pokemon type and pokemon id. Before press ENTER")
			fmt.Scanf("%s %s %s", pokemonData[0], pokemonData[1], pokemonData[2])

			err := pokedex.UpdatePokemon(id, pokemonData)

			if err != nil {
				fmt.Println(err.Error())
			}

		}

		if option == 5 {
			var id int

			fmt.Println("Write a pokemon id for delete")
			fmt.Scanf("%d", &id)

			err := pokedex.DeletePokemon(id)

			if err != nil {
				fmt.Println(err.Error())
			}
		}

		if option == 6 {
			fmt.Println("\n\nBye...")
			break

		} else {
			fmt.Println("Invalid option!!!")
		}

	}

	
}

func menu(){
	fmt.Println("\nChoose a option!!!")
	fmt.Println("\n1 - Add a Pokemon")
	fmt.Println("2 - Show a Pokemon")
	fmt.Println("3 - Show all Pokemons")
	fmt.Println("4 - Update a Pokemon")
	fmt.Println("5 - Delete a Pokemon")
	fmt.Println("6 - Exit")
}

func banner(){
	banner :=
	
	`
	===============================================================
	_______  _______  ___   _  _______  ______   _______  __   __ 
	|       ||       ||   | | ||       ||      | |       ||  |_|  |
	|    _  ||   _   ||   |_| ||    ___||  _    ||    ___||       |
	|   |_| ||  | |  ||      _||   |___ | | |   ||   |___ |       |
	|    ___||  |_|  ||     |_ |    ___|| |_|   ||    ___| |     | 
	|   |    |       ||    _  ||   |___ |       ||   |___ |   _   |
	|___|    |_______||___| |_||_______||______| |_______||__| |__|

	==============================================================
											  
	`

	fmt.Println(banner)
}