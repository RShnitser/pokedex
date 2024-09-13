package main

import(
	"fmt"
	"pokedex/internal/pokeapi"
)

func (state *appState)commandExplore(args ...string) error{
	
	if len(args) != 1{
		return fmt.Errorf("Provide a location name")
	}
	name := args[0]
	data, err := pokeapi.GetLocation("https://pokeapi.co/api/v2/location-area/" + name, &state.cache)
	if err != nil{
		return err
	}

	fmt.Printf("Exploring %s...\n", name)
	fmt.Println("Found pokemon:")
	for _, pokemon := range data.PokemonEncounters{
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	
	return nil
}