package main

import(
	"fmt"
)

func (state *appState)commandPokedex(args ...string) error{
	fmt.Println("Your Pokedex:")
	for k, _ := range state.pokedex{
		fmt.Printf(" -  %s\n", k)
	}
	return nil
}