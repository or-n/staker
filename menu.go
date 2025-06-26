package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
)

type State int

const (
	StateMenu State = iota
	StateGame
	StateOptions
	StateExit
)

var (
	SimulationState State
	buttonSize      = NewVector2(200, 100)
)

func MenuDraw() {
	x := (WindowSize.X - buttonSize.X) * 0.5
	y := (WindowSize.Y - buttonSize.Y*4) * 0.5
	rect := NewRectangle(x, y, buttonSize.X, buttonSize.Y)
	if clicked(rect, Lang[Play]) {
		SimulationState = StateGame
	}
	rect.Y += buttonSize.Y
	if clicked(rect, Lang[Restart]) {
		SimulationState = StateGame
		MainAccount.Balance = 1000
		MainAccount.decided = false
		MainAccount.input = ""
		EventNew()
	}
	rect.Y += buttonSize.Y
	if clicked(rect, Lang[Options]) {
		SimulationState = StateOptions
	}
	rect.Y += buttonSize.Y
	if clicked(rect, Lang[Exit]) {
		SimulationState = StateExit
	}
}
