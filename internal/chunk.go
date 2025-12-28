package internal

import (
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Chunk struct {
	X           int
	Y           int
	Tiles       [][]*Tile
	TreesToDraw []*Tree // trees to draw this frame
}

func NewChunk(x, y int) *Chunk {
	c := &Chunk{
		X:           x,
		Y:           y,
		Tiles:       [][]*Tile{},
		TreesToDraw: []*Tree{},
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
				r := rand.IntN(100) + 1
				if r < 2 {
					t = TileTypeSnowBrush
				} else if r >= 3 && r < 4 {
					if rl.Vector2Distance(center, this) > 20*16 {
						t = TileTypeTreeBig
					}
				} else if r >= 4 && r < 6 {
					if rl.Vector2Distance(center, this) > 20*16 {
						t = TileTypeTreeSmall
					}
				} else if r == 10 {
					t = TileTypeTwig1
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

func (chunk *Chunk) Draw(p *Player, tileset *Tileset) {
	chunk.TreesToDraw = make([]*Tree, 0)
	offsetX := chunk.X * 16 * 16
	offsetY := chunk.Y * 16 * 16
	for ty := 0; ty < 16; ty++ {
		for tx := 0; tx < 16; tx++ {
			t := chunk.Tiles[ty][tx]
			if t.Type == TileTypeTreeBig || t.Type == TileTypeTreeSmall {
				chunk.TreesToDraw = append(chunk.TreesToDraw, NewTree(offsetX+tx*16, offsetY+ty*16))
			}
			t.Draw(tileset, offsetX+tx*16, offsetY+ty*16)

			// player collision
			dist := rl.Vector2Distance(rl.NewVector2(p.X+8, p.Y+16+8), rl.NewVector2(float32(offsetX+tx*16+8), float32(offsetY+ty*16+8)))
			if dist < 8 {
				t.PlayerSteppedOn = true
			}
		}
	}
}
