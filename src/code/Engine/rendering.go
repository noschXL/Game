package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func ResetWindow() {
	rl.ClearBackground(rl.DarkGray)
}

func BeginDrawing() {
	rl.BeginDrawing()
	ResetWindow()
}

func EndDrawing() {
	rl.EndDrawing()
}
