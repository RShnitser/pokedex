package main

func (state *appState)commandExit() error{
	state.running = false
	return nil
}