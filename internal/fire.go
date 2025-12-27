package internal

import rl "github.com/gen2brain/raylib-go/raylib"

type Fire struct {
	Tileset      *Tileset
	Health       int
	Frame        int
	FrameSpeed   int
	FrameCounter float32
}

func NewFire() *Fire {
	return &Fire{
		Tileset:    NewTileset("./assets/CampFireFinished.png", 64, 64),
		Health:     100,
		FrameSpeed: 2,
	}
}

func (f *Fire) Update() {
	f.FrameCounter += rl.GetFrameTime()
	if f.FrameCounter >= 1/float32(f.FrameSpeed) {
		f.FrameCounter = 0
		f.Frame += 1
		if f.Frame == 5 {
			f.Frame = 0
		}
	}
}

func (f *Fire) Draw() {
	f.Tileset.DrawTile(f.Frame, 0, -32, -64)
}
