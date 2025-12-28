package internal

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	TreeTypeBig uint8 = iota
	TreeTypeSmall
)

type Tree struct {
	X int
	Y int
}

func NewTree(x, y int) *Tree {
	return &Tree{
		X: x,
		Y: y,
	}
}

func (t *Tree) DrawBottom(source rl.Texture2D) {
	s := rl.NewRectangle(0, 5*16, 6*16, 4*16)
	d := rl.NewRectangle(float32(t.X), float32(t.Y+5*16), 6*16, 4*16)
	rl.DrawTexturePro(source, s, d, rl.NewVector2(0, 0), 0, rl.White)
}

func (t *Tree) DrawTop(source rl.Texture2D) {
	s := rl.NewRectangle(0, 0, 6*16, 6*16)
	d := rl.NewRectangle(float32(t.X), float32(t.Y), 6*16, 6*16)
	rl.DrawTexturePro(source, s, d, rl.NewVector2(0, 0), 0, rl.White)
}
