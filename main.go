package main

import (
	"embed"
	. "github.com/gen2brain/raylib-go/raylib"
	"syscall/js"
)

func isTabFocused() bool {
	document := js.Global().Get("document")
	if !document.Truthy() {
		return false
	}
	hasFocus := document.Call("hasFocus")
	return hasFocus.Bool()
}

func viewportSize() (width, height int) {
	window := js.Global().Get("window")
	width = window.Get("innerWidth").Int()
	height = window.Get("innerHeight").Int()
	return
}

var (
	//go:embed asset
	ASSETS  embed.FS
	BgMusic Music
	update  = func() {
		vw, vh := viewportSize()
		WindowSize = NewVector2(float32(vw), float32(vh))
		SetMusicVolume(BgMusic, MusicVolume)
		if isTabFocused() {
			UpdateMusicStream(BgMusic)
		}
		if IsKeyDown(KeyEscape) {
			SimulationState = StateMenu
		}
		BeginDrawing()
		ClearBackground(WindowBg)
		switch SimulationState {
		case StateMenu:
			MenuDraw()
		case StateGame:
			// GifUpdate()
			AccountUpdate(&MainAccount)
			// GifDraw()
			EventDraw()
			AccountDraw(&MainAccount)
		case StateOptions:
			OptionsDraw()
		}
		DrawRectangle(0, 0, 130, 80, guiBg)
		DrawFPS(30, 30)
		EndDrawing()
	}
)

func main() {
	AddFileSystem(ASSETS)
	InitWindow(1920, 1080, "Staker")
	InitAudioDevice()
	BgMusic = LoadMusicStream("asset/TetrisThemeDubstep.ogg")
	PlayMusicStream(BgMusic)
	defer func() {
		// if err := Save(AccountFile, &MainAccount); err != nil {
		// 	fmt.Println("Failed to save account:", err)
		// }
		CloseWindow()
	}()
	SetTargetFPS(600)
	FontInit()
	EventNew()
	AccountInit()
	// GifInit()
	// SetExitKey(0)
	SimulationState = StateGame
	SetMainLoop(update)
	for !WindowShouldClose() && SimulationState != StateExit {
		update()
	}
}
