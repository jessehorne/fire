package internal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	PlayerDirectionUp uint8 = iota
	PlayerDirectionDown
	PlayerDirectionLeft
	PlayerDirectionRight
)

type Player struct {
	X            float32
	Y            float32
	Speed        int // how many pixels a second to move
	Frame        int // which frame to draw
	FrameSpeed   int // how often to change frames a second
	FrameCounter float32
	Direction    uint8
	Source       rl.Texture2D
	WalkingSound rl.Sound
	TwigCount    int
}

func NewPlayer() *Player {
	p := &Player{
		X:            3 * 16,
		Y:            3 * 16,
		Speed:        60,
		Frame:        0,
		FrameSpeed:   4,
		Direction:    PlayerDirectionDown,
		Source:       rl.LoadTexture("./assets/final/run-Sheet.png"),
		WalkingSound: rl.LoadSound("./assets/final/walking.ogg"),
	}
	rl.SetSoundVolume(p.WalkingSound, 0.5)
	rl.SetSoundPitch(p.WalkingSound, 1.5)
	return p
}

func (p *Player) Rectangle() rl.Rectangle {
	return rl.NewRectangle(p.X, p.Y, 16, 32)
}

func (p *Player) Update() {
	// movement
	dir := PlayerDirectionDown
	vel := rl.NewVector2(0, 0)
	if rl.IsKeyDown(rl.KeyW) {
		dir = PlayerDirectionUp
		vel.Y = -1
	} else if rl.IsKeyDown(rl.KeyS) {
		dir = PlayerDirectionDown
		vel.Y = 1
	}
	if rl.IsKeyDown(rl.KeyA) {
		dir = PlayerDirectionLeft
		vel.X = -1
	} else if rl.IsKeyDown(rl.KeyD) {
		dir = PlayerDirectionRight
		vel.X = 1
	}
	if vel.X != 0 || vel.Y != 0 {
		if !rl.IsSoundPlaying(p.WalkingSound) {
			rl.PlaySound(p.WalkingSound)
		}
		norm := rl.Vector2Normalize(vel)
		p.X += norm.X * rl.GetFrameTime() * float32(p.Speed)
		p.Y += norm.Y * rl.GetFrameTime() * float32(p.Speed)
		p.Direction = dir

		p.FrameCounter += rl.GetFrameTime()
		if p.FrameCounter >= 1/float32(p.FrameSpeed) {
			p.Frame += 1
			if p.Frame == 8 {
				p.Frame = 0
			}
			p.FrameCounter = 0
		}
	} else {
		if rl.IsSoundPlaying(p.WalkingSound) {
			rl.PauseSound(p.WalkingSound)
		}
	}
}

func (p *Player) Draw() {
	startX := 0
	w := float32(20)
	h := float32(20)
	if p.Direction == PlayerDirectionUp {
		startX = 16
	} else if p.Direction == PlayerDirectionDown {
		startX = 8
	} else if p.Direction == PlayerDirectionLeft {
		w = -w
	}

	s := rl.NewRectangle(float32(startX*80+p.Frame*80+30), 30, w, 20)
	d := rl.NewRectangle(p.X, p.Y, 20, h)
	rl.DrawTexturePro(p.Source, s, d, rl.NewVector2(0, 0), 0, rl.White)
}
