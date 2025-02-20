package engine

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const PLAYERID int32 = 0

var Entitys []Entity

type Entity struct {
	Position Vector2D
	Velocity Vector2D
	SpriteID int32
	ID       int32
}

func NewPlayer(x, y float32, spriteID int32) Entity {
	return Entity{
		Position: Vector2D{x, y},
		Velocity: Vector2D{},
		SpriteID: spriteID,
		ID:       0,
	}
}

func UpdatePlayer(playerEntity *Entity, deltatime float32) {
	accelerationVec := Vector2D{0, 0}
	directions := 0

	if rl.IsGamepadAvailable(0) {
		x := math.Round(float64(rl.GetGamepadAxisMovement(0, 0) * 100))
		y := math.Round(float64(rl.GetGamepadAxisMovement(0, 1) * 100))
		force := float32(math.Min(1.5, math.Sqrt(x*x+y*y)/50))
		println(force)
		angle := math.Atan2(y, x)

		accelerationVec.X = float32(math.Cos(angle)) * force
		accelerationVec.Y = float32(math.Sin(angle)) * force
	} else {

		if rl.IsKeyDown(rl.KeyW) {
			accelerationVec.Y -= 1.5 * (deltatime / (1 / 120.0))
			directions += 1
		} else if rl.IsKeyDown(rl.KeyS) {
			accelerationVec.Y += 1.5 * (deltatime / (1 / 120.0))
			directions += 1
		}
		if rl.IsKeyDown(rl.KeyA) {
			accelerationVec.X -= 1.5 * (deltatime / (1 / 120.0))
			directions += 1
		} else if rl.IsKeyDown(rl.KeyD) {
			accelerationVec.X += 1.5 * (deltatime / (1 / 120.0))
			directions += 1
		}

		if directions == 2 {
			accelerationVec.X *= (1 / math.Sqrt2)
			accelerationVec.Y *= (1 / math.Sqrt2)
		}
	}
	playerEntity.Position.X += accelerationVec.X * (deltatime / (1 / 120.0))
	playerEntity.Position.Y += accelerationVec.Y * (deltatime / (1 / 120.0))

	playerEntity.Velocity.Y *= 0.5
	playerEntity.Velocity.X *= 0.5
}

func UpdateEntitys() {}
