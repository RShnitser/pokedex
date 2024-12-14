package main

import "fmt"

func (state *appState)commandExit(args ...string) error{
	state.running = false
	fmt.Println("Closing the Pokedex... Goodbye!")
	return nil
}