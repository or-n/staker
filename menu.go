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
	button          = NewVector2(200, 100)
)

func clicked(rec Rectangle, label string) bool {
	pad := f32(4)
	rec.X += pad
	rec.Y += pad
	rec.Width -= pad * 2
	rec.Height -= pad * 2
	DrawRectangleRec(rec, guiBg)
	measure := MeasureTextEx(MainFont, label, 20, 2)
	space := NewVector2(rec.Width-measure.X, rec.Height-measure.Y)
	start := Vector2Scale(space, 0.5)
	p := Vector2Add(NewVector2(rec.X, rec.Y), start)
	DrawTextEx(MainFont, label, p, 20, 2, White)
	if !IsMouseButtonPressed(MouseButtonLeft) {
		return false
	}
	cursor := GetMousePosition()
	return CheckCollisionPointRec(cursor, rec)
}

func MenuDraw() {
	x := (WindowSize.X - button.X) * 0.5
	y := (WindowSize.Y - button.Y*4) * 0.5
	rect := NewRectangle(x, y, button.X, button.Y)
	if clicked(rect, Lang[Play]) {
		SimulationState = StateGame
	}
	rect.Y += button.Y
	if clicked(rect, Lang[Restart]) {
		SimulationState = StateGame
		MainAccount.Balance = 1000
		MainAccount.decided = false
		MainAccount.input = ""
		EventNew()
	}
	rect.Y += button.Y
	if clicked(rect, Lang[Options]) {
		SimulationState = StateOptions
	}
	rect.Y += button.Y
	if clicked(rect, Lang[Exit]) {
		SimulationState = StateExit
	}
}
