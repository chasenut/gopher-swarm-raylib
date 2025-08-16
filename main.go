package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	windowWidth int32 = 1200
	windowHeight int32 = 900
	FPS int32 = 60
)

var (
	running bool = true
	dt float32 = rl.GetFrameTime()
	
	bkgColor rl.Color = rl.NewColor(22, 22, 35, 255)

	hueTint rl.Color
	
	perlinImg *rl.Image 
	perlinTexture rl.Texture2D
)

func input() {

}

func update() {
	running = !rl.WindowShouldClose()
	dt = rl.GetFrameTime()

	hueTint = rl.ColorFromHSV(float32(rl.GetTime() * 10000.0 / 360.0), 1, 1)
}

func drawScene() {
	rl.BeginDrawing()

	rl.ClearBackground(bkgColor)


	rl.DrawTexture(perlinTexture, 0, 0, hueTint)

	rl.DrawText(fmt.Sprintf("%v", rl.GetFPS()), 0, 0, 40, rl.LightGray)
	rl.EndDrawing()


}

func render() {
	
}

func initialize() {
	rl.InitWindow(windowWidth, windowHeight, "hello raylib")
	rl.SetTargetFPS(FPS)

	perlinImg = rl.GenImagePerlinNoise(800, 800, 0, 0, 5)
	perlinTexture = rl.LoadTextureFromImage(perlinImg)


}

func exit() {
	defer rl.CloseWindow()

}

func main() {
	initialize()


	for running {
		input()
		update()
		drawScene()
	}
}
