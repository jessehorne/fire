package internal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Map    *Map
	Player *Player
	Camera rl.Camera2D
}

func NewGame() *Game {
	m := NewMap()
	p := NewPlayer()
	return &Game{
		Map:    m,
		Player: p,
		Camera: rl.Camera2D{
			Target:   rl.NewVector2(p.X, p.Y),
			Offset:   rl.NewVector2(float32(rl.GetScreenWidth())/2, float32(rl.GetScreenHeight())/2),
			Rotation: 0.0,
			Zoom:     2.0,
		},
	}
}

func (g *Game) Update() {
	g.Map.Update()
	g.Player.Update()

	// update camera
	g.Camera.Target = rl.NewVector2(g.Player.X+8, g.Player.Y+8)

	// handle camera zoom
	scroll := rl.GetMouseWheelMove()
	if scroll != 0 {
		if scroll < 0 {
			g.Camera.Zoom -= 0.2
			if g.Camera.Zoom < 1 {
				g.Camera.Zoom = 1
			}
		} else {
			g.Camera.Zoom += 0.2
			if g.Camera.Zoom > 5.2 {
				g.Camera.Zoom = 5.2
			}
		}
	}

	// update map chunk x and y
	g.Map.StartChunkX = int(g.Player.X/16/16) - 2
	g.Map.StartChunkY = int(g.Player.Y/16/16) - 2
}

func (g *Game) Draw() {
	rl.BeginMode2D(g.Camera)
	g.Map.Draw()
	g.Player.Draw()
	rl.EndMode2D()
}
