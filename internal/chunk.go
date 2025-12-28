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
	Twigs       []*Twig
}

func NewChunk(x, y int, p *Player) *Chunk {
	c := &Chunk{
		X:           x,
		Y:           y,
		Tiles:       [][]*Tile{},
		TreesToDraw: []*Tree{},
	}
	c.Generate(p)
	return c
}

func (chunk *Chunk) Generate(p *Player) {
	for y := 0; y < 16; y++ {
		chunk.Tiles = append(chunk.Tiles, []*Tile{})
		for x := 0; x < 16; x++ {
			t := TileTypeSnowFloor

			center := rl.NewVector2(0, 0)
			this := rl.NewVector2(float32(chunk.X*16*16+x*16+8), float32(chunk.Y*16*16+y*16+8))

			corner := 0
			if rl.Vector2Distance(center, this) < 4*16 {
				t = TileTypeDirtFloor

				topLeft := (x == 12 && y == 14) || (x == 13 && y == 13) || (x == 14 && y == 12)
				bottomLeft := (x == 12 && y == 1) || (x == 13 && y == 2) || (x == 14 && y == 3)
				topRight := (x == 1 && y == 12) || (x == 2 && y == 13) || (x == 3 && y == 14)
				bottomRight := (x == 1 && y == 3) || (x == 2 && y == 2) || (x == 3 && y == 1)
				if topLeft {
					corner = 1
				} else if bottomLeft {
					corner = 2
				} else if bottomRight {
					corner = 3
				} else if topRight {
					corner = 4
				}

				r := rand.IntN(10) + 1
				if r == 5 {
					t = TileTypeDirtFloor1
				} else if r == 6 {
					t = TileTypeDirtFloor2
				} else if r == 7 {
					t = TileTypeDirtFloor3
				}
			} else {
				r := rand.IntN(100) + 1
				if r >= 3 && r < 4 {
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

			if t == TileTypeTwig1 {
				chunk.Twigs = append(chunk.Twigs, NewTwig(chunk.X*16*16+x*16, chunk.Y*16*16+y*16, p))
			}
			chunk.Tiles[y] = append(chunk.Tiles[y], &Tile{
				Type:   t,
				Corner: corner,
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
			dist := rl.Vector2Distance(rl.NewVector2(p.X+8, p.Y+16), rl.NewVector2(float32(offsetX+tx*16+8), float32(offsetY+ty*16+8)))
			if dist < 8 {
				t.PlayerSteppedOn = true
			}
		}
	}

	// draw twigs
	deleteCount := 0
	for _, t := range chunk.Twigs {
		t.Update()
		t.Draw()
		if t.Delete {
			deleteCount++
		}
	}

	for i := 0; i < deleteCount; i++ {
		deleteI := -1
		for ii, tt := range chunk.Twigs {
			if tt.Delete {
				deleteI = ii
			}
		}
		if deleteI != -1 {
			chunk.Twigs = append(chunk.Twigs[:deleteI], chunk.Twigs[deleteI+1:]...)
		}
	}
}
