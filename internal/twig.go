package internal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Twig struct {
	X      int
	Y      int
	Player *Player
	Delete bool
	Source rl.Texture2D
	Sound  rl.Sound
}

func NewTwig(x int, y int, p *Player) *Twig {
	return &Twig{
		X:      x,
		Y:      y,
		Player: p,
		Source: rl.LoadTexture("./assets/final/twigs1.png"),
		Sound:  rl.LoadSound("./assets/final/wood.ogg"),
	}
}

func (t *Twig) Draw() {
	rl.DrawTexture(t.Source, int32(t.X), int32(t.Y), rl.White)
	//rl.DrawRectangle(int32(t.X), int32(t.Y), 16, 16, rl.Red)
}

func (t *Twig) Update() {
	pv := rl.NewVector2(t.Player.X, t.Player.Y)
	tv := rl.NewVector2(float32(t.X), float32(t.Y))
	if rl.Vector2Distance(pv, tv) < 32 {
		t.Delete = true
		t.Player.TwigCount += 1
		rl.PlaySound(t.Sound)
	}
}
