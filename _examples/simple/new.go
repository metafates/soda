package main

func New() *State {
	return &State{
		keyMap: NewKeyMap(),
	}
}
