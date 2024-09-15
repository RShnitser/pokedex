package main

import(
	"fmt"
	"math/rand/v2"
	"pokedex/internal/pokeapi"
)

func (state *appState)commandCatch(args ...string) error{
	if len(args) != 1{
		return fmt.Errorf("Provide a pokemon name")
	}
	name := args[0]

	data, err := pokeapi.GetPokemon(name, &state.cache)
	if err != nil{
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	caught := rand.IntN(636) - data.BaseExperience >= 0

	if caught{
		fmt.Printf("%s was caught!\n", name)
		state.pokedex[data.Name] = data
		fmt.Println("You may now inspect it with the inspect command.")
	}else{
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}