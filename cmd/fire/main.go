package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"github.com/jessehorne/fire/internal"
)

func main() {
	rl.InitWindow(800, 450, "Fire | A game by Quite Frankly and jessefromgeorgia")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	g := internal.NewGame()

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		
		g.Update()
		g.Draw()

		rl.EndDrawing()
	}
}
