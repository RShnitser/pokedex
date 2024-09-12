package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"pokedex/internal/pokecache"
	"time"
)

type appState struct{
	cmdMap map[string]cliCommand
	running bool
	cfg config
	cache pokecache.Cache
}

type config struct{
	next *string
	previous *string
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func initState()*appState{
	startUrl := "https://pokeapi.co/api/v2/location-area"
	state := appState{}
	state.running = true
	state.cmdMap = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    state.commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    state.commandExit,
		},
		"map": {
			name:        "map",
			description: "Get the next 20 locations",
			callback:    state.commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous 20 locations",
			callback:    state.commandMapB,
		},
	}
	state.cfg = config{
		next: &startUrl,
		previous: nil,
	}
	state.cache = pokecache.NewCache(5 * time.Second)
	return &state
}

func (state *appState)run(){
	scanner := bufio.NewScanner(os.Stdin)
	
	for state.running{
		fmt.Printf("pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		input = strings.ToLower(input)
		words := strings.Fields(input)

		cmd, ok := state.cmdMap[words[0]]
		if !ok{
			fmt.Println("Invalid command")
			continue
		}

		err := cmd.callback()
		if err != nil{
			fmt.Println(err)
			continue
		}
	}
}