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
	}else{
		fmt.Printf("%s escaped!\n", name)
	}

	// fmt.Println(data.Name)
	// fmt.Println(data.BaseExperience)
	// fmt.Println(data.Height)
	// fmt.Println(data.Weight)
	// for _, stat := range data.Stats{
	// 	fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	// }
	// for _, typeData := range data.Types{
	// 	fmt.Printf(" - %s\n", typeData.Type.Name)
	// }


	return nil
}