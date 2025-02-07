package main

import (
	"image/color"
	"math"
	"math/rand"

	noise "github.com/aquilax/go-perlin"
	rl "github.com/gen2brain/raylib-go/raylib"
	engine "github.com/noschXL/Game/code/Engine"
)

func main() {
	winsize := engine.Vector2D{X: 1280, Y: 920}
	engine.Initialize(winsize, "Sole Survivor", 120)
	Gamenoise := noise.NewPerlin(1.75, 2, 8, rand.Int63())
	running := true
	defer rl.CloseWindow()
	engine.BeginDrawing()

	gamemap := rl.LoadRenderTexture(1280, 920)
	rl.BeginTextureMode(gamemap)

	var tries int32 = 1
	valid := false
	for ; tries <= 15 && valid == false; tries++ {
		println("on try: ", tries)
		var whitepixels int32 = 0

		for i := 0.0; i < float64(winsize.X); i++ {
			for j := 0.0; j < float64(winsize.Y); j++ {
				brightness := math.Round(Gamenoise.Noise2D(i/float64(winsize.X), j/float64(winsize.Y)) * 255)
				if brightness < 100 {
					brightness = 0
				} else {
					brightness = 255
					whitepixels++
				}
				rl.DrawPixel(int32(i), int32(j), color.RGBA{uint8(brightness), uint8(brightness), uint8(brightness), 255})
			}
		}
		valid = whitepixels > int32(winsize.X)*int32(winsize.Y)/6
	}
	println(valid)
	rl.EndTextureMode()
	for running == true {
		engine.BeginDrawing()
		rl.DrawTexture(gamemap.Texture, 0, 0, rl.White)
		engine.EndDrawing()
		if rl.WindowShouldClose() {
			running = false
		}
	}
}
