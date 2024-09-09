package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
)

type appState struct{
	cmdMap map[string]cliCommand
	running bool
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func initState()*appState{
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
	}
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