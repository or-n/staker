package main

import (
	"encoding/hex"
	"fmt"
	. "github.com/gen2brain/raylib-go/raylib"
	"log"
)

var (
	en      = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	pl      = []rune("ąćęłńóśźżĄĆĘŁŃÓŚŹŻ")
	symbols = []rune(" .,!?:;_()")
	// codepoints = append(append(en, pl...), symbols...)
	codepoints = en
	MainFont   Font
	fontFile   = "asset/FiraCode-Bold.ttf"
)

func FontInit() {
	// MainFont = LoadFontEx(fontFile, 32, codepoints, i32(len(codepoints)))
	fmt.Println("Loading font")
	data, err := ASSETS.ReadFile(fontFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(data))
	fontData := LoadFontData(data, 20, codepoints, i32(len(codepoints)), 0)
	fmt.Println(fontData)
	// MainFont = LoadFontFromMemory(".ttf", data, 20, []rune{})
	// fmt.Println("Font loaded")
	MainFont = GetFontDefault()
}
