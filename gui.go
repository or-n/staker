package main

import (
	"fmt"
	. "github.com/gen2brain/raylib-go/raylib"
)

var (
	guiBg       = NewColor(0, 0, 0, 195)
	sliderFg    = NewColor(1, 1, 1, 127)
	sliderWidth = float32(10)
	sliderOn    = make(map[string]bool)
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

func slider(rec Rectangle, label string, value, min, max f32) f32 {
	pad := f32(4)
	rec.X += pad
	rec.Y += pad
	rec.Width -= pad * 2
	rec.Height -= pad * 2
	DrawRectangleRec(rec, guiBg)
	labelValue := fmt.Sprintf("%s: %v", label, value)
	measure := MeasureTextEx(MainFont, labelValue, 20, 2)
	space := NewVector2(rec.Width-measure.X, rec.Height-measure.Y)
	start := Vector2Scale(space, 0.5)
	p := Vector2Add(NewVector2(rec.X, rec.Y), start)
	DrawTextEx(MainFont, labelValue, p, 20, 2, White)
	t := (value - min) / (max - min)
	cursor := GetMousePosition()
	if CheckCollisionPointRec(cursor, rec) {
		if IsMouseButtonPressed(MouseButtonLeft) {
			sliderOn[label] = true
		}
	}
	if IsMouseButtonUp(MouseButtonLeft) {
		sliderOn[label] = false
	}
	if value, ok := sliderOn[label]; ok && value {
		t = (cursor.X - rec.X) / rec.Width
		if t < 0 {
			t = 0
		}
		if t > 1 {
			t = 1
		}
	}
	rec2 := rec
	rec2.X = Lerp(rec.X, rec.X+rec.Width-sliderWidth, t)
	rec2.Width = sliderWidth
	DrawRectangleRec(rec2, sliderFg)
	return Lerp(min, max, t)
}
