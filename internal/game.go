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

func (g *Game) HandlePlayerCollisions() {
	// with trees
	ppos := rl.NewVector2(g.Player.X+8, g.Player.Y+16)
	//rl.DrawCircle(int32(ppos.X), int32(ppos.Y), 8, rl.Red)
	for _, t := range g.Map.TreesToDraw {
		tpos := rl.NewVector2(float32(t.X+3*16), float32(t.Y+7*16))
		//rl.DrawCircle(int32(tpos.X), int32(tpos.Y), 8, rl.Blue)
		dist := rl.Vector2Distance(ppos, tpos)
		if dist < 16 {
			if dist == 0 {
				dist = 0.1
			}
			normal := rl.Vector2Normalize(rl.Vector2Subtract(ppos, tpos))
			overlap := 16 - dist
			pushback := rl.Vector2Scale(normal, overlap)
			g.Player.X += pushback.X
			g.Player.Y += pushback.Y
			ppos = rl.Vector2Add(ppos, pushback)
		}
	}
}

func (g *Game) Draw() {
	rl.BeginMode2D(g.Camera)
	g.Map.Draw(g.Player)
	g.Fire.Draw()

	g.HandlePlayerCollisions()

	// draw bottoms of trees
	for _, t := range g.Map.TreesToDraw {
		// check if player behind trees
		pbox := g.Player.Rectangle()
		tbox := rl.NewRectangle(float32(t.X), float32(t.Y), 96, 144-64)
		//rl.DrawRectanglePro(tbox, rl.NewVector2(0, 0), 0.0, rl.Red)
		//rl.DrawRectanglePro(pbox, rl.NewVector2(0, 0), 0.0, rl.Blue)
		if rl.CheckCollisionRecs(tbox, pbox) {
			t.PlayerBehind = true
		}
		t.DrawBottom(g.TreeSource)
	}

	// draw player
	g.Player.Draw()

	// draw tops of trees
	for _, t := range g.Map.TreesToDraw {
		t.DrawTop(g.TreeSource)
		t.PlayerBehind = false
	}

	g.Map.TreesToDraw = nil

	rl.EndMode2D()
}
