package internal

const (
	TileTypeSnowFloor uint8 = iota
	TileTypeDirtFloor
	TileTypeSnowBrush
)

type Tile struct {
	Type uint8
}

func NewTile(t uint8) *Tile {
	return &Tile{t}
}

func (t *Tile) Draw(tileset *Tileset, x, y int) {
	tx := -0
	ty := -1
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
	if t.Type == TileTypeSnowBrush {
		tileset.DrawTile(0, 0, float32(x), float32(y))
	}
	tileset.DrawTile(tx, ty, float32(x), float32(y))
}
