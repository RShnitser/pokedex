package main

func (state *appState)commandExit(args ...string) error{
	state.running = false
	return nil
}