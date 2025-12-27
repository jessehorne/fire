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
	Tileset      *Tileset
}

func NewPlayer() *Player {
	return &Player{
		X:          8 * 16,
		Y:          8 * 16,
		Speed:      100,
		Frame:      0,
		FrameSpeed: 4,
		Direction:  PlayerDirectionDown,
		Tileset:    NewTileset("./assets/character.png", 16, 32),
	}
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
		norm := rl.Vector2Normalize(vel)
		p.X += norm.X * rl.GetFrameTime() * float32(p.Speed)
		p.Y += norm.Y * rl.GetFrameTime() * float32(p.Speed)
		p.Direction = dir

		p.FrameCounter += rl.GetFrameTime()
		if p.FrameCounter >= 1/float32(p.FrameSpeed) {
			p.Frame += 1
			if p.Frame == 4 {
				p.Frame = 0
			}
			p.FrameCounter = 0
		}
	}
}

func (p *Player) Draw() {
	tx := p.Frame
	ty := 0
	if p.Direction == PlayerDirectionUp {
		ty = 2
	} else if p.Direction == PlayerDirectionDown {
		ty = 0
	} else if p.Direction == PlayerDirectionLeft {
		ty = 3
	} else if p.Direction == PlayerDirectionRight {
		ty = 1
	}

	p.Tileset.DrawTile(tx, ty, p.X, p.Y)
}
