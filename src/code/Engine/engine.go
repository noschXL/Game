package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Initialize(winsize Vector2D, title string, FPS int32) {
	rl.InitWindow(int32(winsize.X), int32(winsize.Y), title)

	rl.SetTargetFPS(FPS)
}

func DrawImage(x int32, y int32, img *rl.Image) {
	rl.DrawTexture(rl.LoadTextureFromImage(img), x, y, rl.White)
}

func Update() {
	UpdateTiles()
	UpdateEntitys()
}
