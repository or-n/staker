package main

import (
	"fmt"
	. "github.com/gen2brain/raylib-go/raylib"
	"math"
)

var (
	MusicVolume = f32(0.5)
	ShowFPS     = false
	h           = f32(50)
	spacing     = h
)

func bool2string(x bool) string {
	if x {
		return "true"
	}
	return "false"
}

func OptionsDraw() {
	w := f32(400)
	x := (WindowSize.X - w) * 0.5
	y := (WindowSize.Y - h*6) * 0.5
	rect := NewRectangle(x, y, w, h)
	MusicVolume = slider(rect, Lang[Volume], MusicVolume, 0, 1)
	rect.Y += spacing
	labelFPS := fmt.Sprintf("show FPS: %s", bool2string(ShowFPS))
	if clicked(rect, labelFPS) {
		ShowFPS = !ShowFPS
	}
	rect.Y += spacing
	labelOdd0 := fmt.Sprintf("show odd 0: %s", bool2string(ShowOdd0))
	if clicked(rect, labelOdd0) {
		ShowOdd0 = !ShowOdd0
	}
	rect.Y += spacing
	labelWinOptions := Lang[WinOptions]
	next_n = int32(math.Round(f64(slider(rect, labelWinOptions, float32(next_n), 1, 5))))
	rect.Y += spacing
	labelMinValue := fmt.Sprintf("min %s", Lang[Value])
	min_value = slider(rect, labelMinValue, min_value, 0, 2)
	rect.Y += spacing
	labelMaxValue := fmt.Sprintf("max %s", Lang[Value])
	max_value = slider(rect, labelMaxValue, max_value, 0, 2)
	if min_value > max_value {
		min := max_value
		max_value = min_value
		min_value = min
	}
}
