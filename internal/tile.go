package internal

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	TileTypeSnowFloor uint8 = iota
	TileTypeDirtFloor
	TileTypeDirtFloor1
	TileTypeDirtFloor2
	TileTypeDirtFloor3
	TileTypeSnowBrush
	TileTypeTreeBig
	TileTypeTreeSmall
	TileTypeTwig1
	TileTypeTwig2
	TileTypeTwig3
)

type Tile struct {
	Type            uint8
	PlayerSteppedOn bool
	PlayerStepTime  float32
	Corner          int // if 0, don't draw a corner
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

	tx := -1
	ty := -1
	switch t.Type {
	case TileTypeSnowFloor:
		tx = 1
		ty = 1
	case TileTypeDirtFloor:
		tx = 1
		ty = 3
	case TileTypeDirtFloor1:
		tx = 0
		ty = 3
	case TileTypeDirtFloor2:
		tx = 0
		ty = 4
	case TileTypeDirtFloor3:
		tx = 1
		ty = 4
	}

	if tx != -1 {
		tileset.DrawTile(tx, ty, float32(x), float32(y))
	} else {
		tileset.DrawTile(1, 1, float32(x), float32(y))
	}
	if t.PlayerSteppedOn && t.Type == TileTypeSnowFloor {
		tileset.DrawTile(4, 4, float32(x), float32(y))
	}

	if t.Corner != 0 {
		switch t.Corner {
		case 1:
			tileset.DrawTile(2, 1, float32(x), float32(y))
			tileset.DrawTile(1, 2, float32(x), float32(y))
			tileset.DrawTile(2, 2, float32(x), float32(y))
		case 2:
			tileset.DrawTile(2, 0, float32(x), float32(y))
			tileset.DrawTile(1, 0, float32(x), float32(y))
			tileset.DrawTile(2, 1, float32(x), float32(y))
		case 3:
			tileset.DrawTile(0, 0, float32(x), float32(y))
			tileset.DrawTile(1, 0, float32(x), float32(y))
			tileset.DrawTile(0, 1, float32(x), float32(y))
		case 4:
			tileset.DrawTile(0, 2, float32(x), float32(y))
			tileset.DrawTile(0, 1, float32(x), float32(y))
			tileset.DrawTile(1, 2, float32(x), float32(y))
		}
	}
}
