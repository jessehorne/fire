package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"github.com/jessehorne/fire/internal"
)

func main() {
	rl.InitWindow(800, 450, "Fire | A game by Claragraal, Quite Frankly and jessefromgeorgia")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()
	for !rl.IsAudioDeviceReady() {
	}

	bgm := rl.LoadMusicStream("./assets/final/music.ogg")
	defer rl.UnloadMusicStream(bgm)
	bgm.Looping = true
	rl.PlayMusicStream(bgm)

	g := internal.NewGame()
	lightMask := rl.LoadRenderTexture(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()))

	for !rl.WindowShouldClose() {
		rl.UpdateMusicStream(bgm)
		// lighting stuff
		rl.BeginTextureMode(lightMask)
		rl.ClearBackground(rl.Blank)
		rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.NewColor(0, 0, 0, 240))
		rl.BeginBlendMode(rl.BlendAlpha)
		firePos := rl.GetWorldToScreen2D(rl.NewVector2(0, 0), g.Camera)
		a := uint8(g.Fire.Stage * 30)
		rl.DrawCircleGradient(int32(firePos.X), int32(firePos.Y-16), 255*g.Camera.Zoom, rl.NewColor(255, 150, 0, a), rl.Blank)
		rl.EndBlendMode()
		rl.EndTextureMode()

		// normal drawing stuff
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		g.Update()
		g.Draw()

		// draw light mask
		s := rl.NewRectangle(0, 0, float32(lightMask.Texture.Width), float32(-lightMask.Texture.Height))
		d := rl.NewRectangle(0, 0, float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight()))
		rl.DrawTexturePro(lightMask.Texture, s, d, rl.NewVector2(0, 0), 0.0, rl.White)

		rl.EndDrawing()
	}
}
