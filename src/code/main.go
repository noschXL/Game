package main

import (
	"image/color"
	"math"
	"math/rand"

	noise "github.com/aquilax/go-perlin"
	rl "github.com/gen2brain/raylib-go/raylib"
	engine "github.com/noschXL/Game/code/Engine"
)

func generateMap(seed int64, mapsize engine.Vector2D, gamemap rl.RenderTexture2D) (rl.RenderTexture2D, int32) {

	Gamenoise := noise.NewPerlin(1.8, 2, 8, seed)
	rl.BeginTextureMode(gamemap)
	var whitepixels int32 = 0
	for i := 0.0; i < float64(mapsize.X); i++ {
		for j := 0.0; j < float64(mapsize.Y); j++ {
			brightness := math.Round(Gamenoise.Noise2D(i/float64(mapsize.X), j/float64(mapsize.Y)) * 255)
			if brightness < 100 {
				brightness = 0
			} else {
				brightness = 255
				whitepixels += 1
			}
			rl.DrawPixel(int32(i), int32(j), color.RGBA{uint8(brightness), uint8(brightness), uint8(brightness), 255})
		}
	}

	return gamemap, whitepixels
}

func main() {
	winsize := engine.Vector2D{X: 1280, Y: 920}
	engine.Initialize(winsize, "Sole Survivor", 120)
	running := true
	defer rl.CloseWindow()
	engine.BeginDrawing()

	var tries int32 = 1
	var valid bool = false
	var whitepixel int32 = 0

	minwhitepx := int32(winsize.X) * int32(winsize.Y) / 8

	var gamemap rl.RenderTexture2D = rl.LoadRenderTexture(1280, 920)
	var seed int64 = 0

	for ; tries <= 30 && valid == false; tries++ {
		whitepixel = 0
		seed = rand.Int63()
		println("on try: ", tries)
		gamemap, whitepixel = generateMap(seed, winsize, gamemap)
		valid = int32(whitepixel) > minwhitepx
	}

	println("seed for this map: ", seed)
	rl.EndTextureMode()
	player := engine.NewPlayer(winsize.X/2, winsize.Y/2, 0)
	sprite := engine.LoadPlayerSprite(engine.AppendFilePath(engine.GetPath(), []string{"data", "img", "Entitys", "Player", "player.png"}))

	for running {
		deltatime := rl.GetFrameTime()
		// Update
		engine.UpdatePlayer(&player, deltatime)

		// Draw
		engine.BeginDrawing()
		rl.DrawTexture(gamemap.Texture, 0, 0, rl.White)
		rl.DrawRectangle(int32(player.Position.X), int32(player.Position.Y), 50, 50, rl.Green)
		rl.DrawTexture(sprite, int32(player.Position.X), int32(player.Position.Y), rl.Green)
		engine.EndDrawing()

		//Other Logic
		if rl.WindowShouldClose() {
			running = false
		}
	}
}
