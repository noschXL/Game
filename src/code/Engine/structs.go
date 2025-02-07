package engine

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
	highest := max(vec.X, vec.Y)
	vec.X /= highest
	vec.Y /= highest
}

type Tile struct {
	id       int8
	data     int16
	explored bool // if its a structure and explored
}

type Chunk struct { //a 32 x 32 Area of Tiles
	Tiledata [32 * 32]Tile // 8 bits of identifyers so 256 tiles and 24 bits of data for storage
	X        int32
	Y        int32
}

type Map struct {
	chunks []Chunk
}
