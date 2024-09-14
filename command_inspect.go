package main

import(
	"fmt"
)

func (state *appState)commandInspect(args ...string) error{
	if len(args) != 1{
		return fmt.Errorf("Provide a pokemon name")
	}
	name := args[0]

	data, ok := state.pokedex[name]
	if ok{
		fmt.Printf("Name: %s\n", data.Name)
		fmt.Printf("Height: %v\n", data.Height)
		fmt.Printf("Weight: %v\n", data.Weight)
		fmt.Println("Stats:")
		for _, stat := range data.Stats{
			fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, typeData := range data.Types{
			fmt.Printf(" - %s\n", typeData.Type.Name)
		}
	}else{
		fmt.Printf("%s has not been caught\n", name)
	}

	return nil
}