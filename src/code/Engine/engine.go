package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const PlayerShouldDraw bool = false

func Initialize(winsize Vector2D, title string, FPS int32) {
	rl.InitWindow(int32(winsize.X), int32(winsize.Y), title)

	rl.SetTargetFPS(FPS)
}

func LoadPlayerSprite(path string) rl.Texture2D {
	Texture := rl.LoadTexture(path)
	return Texture
}

// 1: resting
// 2: walking right
// 3: walking left
// 4: walking up
// 5: walking down

func GetPlayerSprite(path string, ID int32) {

}

func DrawImage(x int32, y int32, img *rl.Image) {
	rl.DrawTexture(rl.LoadTextureFromImage(img), x, y, rl.White)
}

func Update(playerentity *Entity, dt float32) {
	UpdateTiles()
	UpdateEntitys()
	UpdatePlayer(playerentity, dt)
}
