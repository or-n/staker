package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
)

var (
	BgMusic Music
	update  = func() {
		vw, vh := viewportSize()
		SetWindowSize(vw, vh)
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
		if ShowFPS {
			DrawRectangle(0, 0, 130, 80, guiBg)
			DrawFPS(30, 30)
		}
		EndDrawing()
	}
)

func main() {
	AddFileSystem(ASSETS)
	SetConfigFlags(FlagVsyncHint | FlagWindowResizable)
	InitWindow(0, 0, "Staker")
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
