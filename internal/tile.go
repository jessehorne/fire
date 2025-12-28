package internal

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	TileTypeSnowFloor uint8 = iota
	TileTypeDirtFloor
	TileTypeSnowBrush
	TileTypeTreeBig
	TileTypeTreeSmall
	TileTypeTwig1
)

type Tile struct {
	Type            uint8
	PlayerSteppedOn bool
	PlayerStepTime  float32
}

func NewTile(t uint8) *Tile {
	return &Tile{
		Type: t,
	}
}

func (t *Tile) Draw(tileset *Tileset, x, y int) {
	// footprints
	if t.PlayerSteppedOn {
		t.PlayerStepTime -= rl.GetFrameTime()
		if t.PlayerStepTime <= 0 {
			t.PlayerStepTime = 10
			t.PlayerSteppedOn = false
		}
	}

	tx := 0
	ty := 0
	switch t.Type {
	case TileTypeSnowFloor:
		tx = 0
		ty = 0
	case TileTypeDirtFloor:
		tx = 10
		ty = 3
	case TileTypeSnowBrush:
		tx = 18
		ty = 2
	}

	tileset.DrawTile(0, 0, float32(x), float32(y))
	tileset.DrawTile(tx, ty, float32(x), float32(y))
	if t.PlayerSteppedOn && t.Type != TileTypeDirtFloor {
		tileset.DrawTile(12, 6, float32(x), float32(y))
	}
}
