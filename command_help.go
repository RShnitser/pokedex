package main

import(
	"fmt"
)

func (state *appState)commandHelp(args ...string) error{
	fmt.Println()
	fmt.Println("Welcome to the pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, v := range state.cmdMap{
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	fmt.Println()
	return nil
}