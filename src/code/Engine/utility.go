package engine

import (
	"math"
	"os"
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Rectangle struct {
	X      int32
	Y      int32
	Width  int32
	Height int32
}

func (rect Rectangle) colliding(other Rectangle) bool {
	return rect.X < other.X+other.Width && rect.X+rect.Width > other.X && rect.Y < other.Y+other.Height && rect.Y+rect.Height > other.Y
}

type Vector2D struct {
	X float32
	Y float32
}

func (vec Vector2D) Add(other Vector2D) Vector2D {
	vec.X += other.X
	vec.Y += other.Y
	return vec
}

func (vec *Vector2D) Normalize() {
	highest := max(
		float32(math.Abs(float64(vec.X))),
		float32(math.Abs(float64(vec.Y))),
	)
	vec.X /= highest
	vec.Y /= highest
}

func (vec Vector2D) Magnitude() float32 {
	value := math.Sqrt(float64(vec.X)*float64(vec.X) + float64(vec.Y)*float64(vec.Y))
	return float32(value)
}

type Tile struct {
	id       int8
	data     int16
	explored bool // if its a structure and explored
	Texture  rl.Texture2D
}

type Chunk struct { //a 32 x 32 Area of Tiles
	Tiledata [32 * 32]Tile // 8 bits of identifyers so 256 tiles and 24 bits of data for storage
	X        int32
	Y        int32
	Texture  rl.Texture2D
}

type Map struct {
	chunks []Chunk
}

func GetPath() string {
	filename, err := os.Executable()
	if err != nil {
		return ""
	}

	filename = filepath.Dir(filename)
	return filename
}

func AppendFilePath(path string, toAppend []string) string {
	for _, appendage := range toAppend {
		path += string(os.PathSeparator) + appendage
	}
	return path
}
