package internal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Fire struct {
	Source       rl.Texture2D
	Stage        int // 0 means out, 1 means low, 2 means med, 3 means full
	StageCounter float32
	Frame        int
	FrameSpeed   int
	FrameCounter float32
}

func NewFire() *Fire {
	return &Fire{
		Source:     rl.LoadTexture("./assets/final/fire.png"),
		Stage:      3,
		FrameSpeed: 3,
	}
}

func (f *Fire) Update() {
	f.FrameCounter += rl.GetFrameTime()
	if f.FrameCounter >= 1/float32(f.FrameSpeed) {
		f.FrameCounter = 0
		f.Frame += 1
		if f.Frame == 3 {
			f.Frame = 0
		}
	}

	if f.Stage != 0 {
		f.StageCounter += rl.GetFrameTime()
	}
	if f.StageCounter >= 5 {
		f.StageCounter = 0
		f.Stage -= 1
		if f.Stage < 0 {
			f.Stage = 0
		}
	}
}

func (f *Fire) Draw() {
	startX := 0
	if f.Stage == 2 {
		startX = 3 * 32
	} else if f.Stage == 1 {
		startX = 6 * 32
	} else if f.Stage == 0 {
		startX = 9 * 32
		f.Frame = 0
	}

	s := rl.NewRectangle(float32(startX+f.Frame*32), 0, 32, 32)
	d := rl.NewRectangle(-16, -16-8, 32, 32)
	rl.DrawTexturePro(f.Source, s, d, rl.NewVector2(0, 0), 0, rl.White)
}
