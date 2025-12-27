package internal

import rl "github.com/gen2brain/raylib-go/raylib"

type Tileset struct {
	Source rl.Texture2D
	TileW  int
	TileH  int
}

func NewTileset(path string, w, h int) *Tileset {
	return &Tileset{
		Source: rl.LoadTexture(path),
		TileW:  w,
		TileH:  h,
	}
}

func (t *Tileset) DrawTile(sx, sy int, dx, dy float32) {
	s := rl.NewRectangle(float32(sx*t.TileW), float32(sy*t.TileH), float32(t.TileW), float32(t.TileH))
	d := rl.NewRectangle(dx, dy, float32(t.TileW), float32(t.TileH))
	rl.DrawTexturePro(t.Source, s, d, rl.NewVector2(0, 0), 0, rl.White)
}
