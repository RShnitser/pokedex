package main

import(
	"fmt"
)

func (state *appState)commandHelp() error{
	fmt.Println()
	fmt.Println("Welcome to the pokedex!")
	fmt.Println("Usage:\n")

	for _, v := range state.cmdMap{
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	fmt.Println()
	return nil
}