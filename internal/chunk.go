package internal

import (
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Chunk struct {
	X     int
	Y     int
	Tiles [][]*Tile
}

func NewChunk(x, y int) *Chunk {
	c := &Chunk{
		X:     x,
		Y:     y,
		Tiles: [][]*Tile{},
	}
	c.Generate()
	return c
}

func (chunk *Chunk) Generate() {
	for y := 0; y < 16; y++ {
		chunk.Tiles = append(chunk.Tiles, []*Tile{})
		for x := 0; x < 16; x++ {
			t := TileTypeSnowBrush

			center := rl.NewVector2(0, 0)
			this := rl.NewVector2(float32(chunk.X*16*16+x*16+8), float32(chunk.Y*16*16+y*16+8))

			if rl.Vector2Distance(center, this) < 7*16 {
				t = TileTypeDirtFloor
			} else {
				if rand.IntN(100)+1 < 2 {
					t = TileTypeSnowBrush
				} else {
					t = TileTypeSnowFloor
				}
			}

			chunk.Tiles[y] = append(chunk.Tiles[y], &Tile{
				Type: t,
			})
		}
	}
}

func (chunk *Chunk) Draw(tileset *Tileset) {
	offsetX := chunk.X * 16 * 16
	offsetY := chunk.Y * 16 * 16
	for ty := 0; ty < 16; ty++ {
		for tx := 0; tx < 16; tx++ {
			chunk.Tiles[ty][tx].Draw(tileset, offsetX+tx*16, offsetY+ty*16)
		}
	}
}
