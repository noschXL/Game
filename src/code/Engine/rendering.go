package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (player Entity) Draw(sprite rl.Texture2D) {
	rl.DrawTexture(sprite, int32(player.Position.X), int32(player.Position.Y), rl.White)
}

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
