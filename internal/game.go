package internal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Map        *Map
	Player     *Player
	Fire       *Fire
	Camera     rl.Camera2D
	TreeSource rl.Texture2D
}

func NewGame() *Game {
	m := NewMap()
	p := NewPlayer()
	f := NewFire()
	return &Game{
		Map:    m,
		Player: p,
		Fire:   f,
		Camera: rl.Camera2D{
			Target:   rl.NewVector2(p.X, p.Y),
			Offset:   rl.NewVector2(float32(rl.GetScreenWidth())/2, float32(rl.GetScreenHeight())/2),
			Rotation: 0.0,
			Zoom:     2.0,
		},
		TreeSource: rl.LoadTexture("./assets/TilesB.png"),
	}
}

func (g *Game) Update() {
	g.Map.Update()
	g.Player.Update()
	g.Fire.Update()

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
	g.Map.StartChunkX = int(g.Player.X/16/16) - 3
	g.Map.StartChunkY = int(g.Player.Y/16/16) - 2
}

func (g *Game) Draw() {
	rl.BeginMode2D(g.Camera)
	g.Map.Draw()
	g.Fire.Draw()

	// draw bottoms of trees
	for _, t := range g.Map.TreesToDraw {
		t.DrawBottom(g.TreeSource)
	}

	// draw player
	g.Player.Draw()

	// draw tops of trees
	for _, t := range g.Map.TreesToDraw {
		t.DrawTop(g.TreeSource)
	}

	g.Map.TreesToDraw = nil

	rl.EndMode2D()
}
