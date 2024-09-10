package main

import(
	"fmt"
	"pokedex/internal/pokeapi"
)

func (state *appState)commandMap() error{
	
	data, err := pokeapi.GetLocations(state.cfg.next)
	if err != nil{
		return err
	}

	state.cfg.previous = data.Previous
	state.cfg.next = data.Next

	for _, loc := range data.Results{
		fmt.Println(loc.Name)
	}

	return nil
}

func (state *appState)commandMapB() error{
	if state.cfg.previous == nil{
		return fmt.Errorf("No previous locations")
	}

	data, err := pokeapi.GetLocations(state.cfg.next)
	if err != nil{
		return err
	}

	state.cfg.previous = data.Previous
	state.cfg.next = data.Next

	for _, loc := range data.Results{
		fmt.Println(loc.Name)
	}

	return nil
}