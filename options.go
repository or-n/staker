package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
)

var (
	MusicVolume = f32(0.5)
)

func slider(rec Rectangle, label string, value, min, max f32) f32 {
	DrawRectangleRec(rec, guiBg)
	p := NewVector2(rec.X, rec.Y)
	DrawTextEx(MainFont, label, p, 20, 2, White)
	return 0.5
}

func OptionsDraw() {
	w := f32(400)
	x := (WindowSize.X - w) * 0.5
	rect := NewRectangle(x, 200, w, 50)
	MusicVolume = slider(rect, Lang[Volume], MusicVolume, 0, 1)
}
