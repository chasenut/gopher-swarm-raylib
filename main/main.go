package main

import (
	"fmt"
	"github.com/chasenut/gorl-learn/gopher"
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth int32 = 1200
	screenHeight int32 = 900
	FPS int32 = 60
)

var (
	running bool = true
	dt float32 = rl.GetFrameTime()

	showDebugPanel bool = true
	
	bkgColor rl.Color = rl.NewColor(22, 22, 35, 255)

	hueTint rl.Color
	
	perlinImg *rl.Image 
	perlinTexture rl.Texture2D

	cameraTarget rl.Vector2
	cameraOffset rl.Vector2
	camera rl.Camera2D
	cameraSpeed float32 = 250
	cameraMoving bool = false
	cameraUp, cameraDown, cameraRight, cameraLeft = false, false, false, false

	gopherTexture rl.Texture2D
	gophers []gopher.Gopher = []gopher.Gopher{}
	gopherCount int32 = 0
	toogleGopherCreation bool = false

)

func createGopher() {
	rect := rl.NewRectangle(
		cameraOffset.X,
		cameraOffset.Y,
		100.0,
		100.0,
		)
	vel := rl.NewVector2(
		(0.5-rand.Float32()) * 50, 
		(0.5-rand.Float32()) * 50,
		)
	gophers = gopher.AddGopher(gophers, rect, vel)
	gopherCount++
}


func input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		cameraMoving = true
		cameraUp = true
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		cameraMoving = true
		cameraDown = true
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		cameraMoving = true
		cameraRight = true
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		cameraMoving = true
		cameraLeft = true
	}
	if rl.IsKeyDown(rl.KeyQ) {
		toogleGopherCreation = true
	}
}

func update() {
	running = !rl.WindowShouldClose()
	dt = rl.GetFrameTime()

	hueTint = rl.ColorFromHSV(float32(rl.GetTime() * 1000 / 36), 1, 1)

	if cameraMoving {
		if cameraUp {
			cameraOffset.Y -= cameraSpeed * dt
		}
		if cameraDown {
			cameraOffset.Y += cameraSpeed * dt
		}
		if cameraRight {
			cameraOffset.X += cameraSpeed * dt
		}
		if cameraLeft {
			cameraOffset.X -= cameraSpeed * dt
		}
	}
	cameraMoving = false
	cameraUp, cameraDown, cameraRight, cameraLeft = false, false, false, false
	camera.Target = cameraOffset

	if toogleGopherCreation {
		for range 1 {
			createGopher()
		}
	}
	toogleGopherCreation = false
	gopher.UpdateGophers(gophers, dt)
}

func drawDebugPanel() {
	if !showDebugPanel {
		return
	}
	rl.DrawText(fmt.Sprintf("FPS: %.3f", float32(1.0 / dt)),
		int32(cameraOffset.X)-screenWidth/2+20, 
		int32(cameraOffset.Y)-screenHeight/2+20, 
		40, rl.LightGray)

	rl.DrawText(fmt.Sprintf("Gophers: %v", gopherCount),
		int32(cameraOffset.X)-screenWidth/2+20, 
		int32(cameraOffset.Y)-screenHeight/2+60, 
		40, rl.LightGray)
}

func drawScene() {
	rl.ClearBackground(bkgColor)


	rl.DrawTexture(perlinTexture, 0, 0, hueTint)

	gopher.DrawGophers(gophers, gopherTexture)

	gopher.DrawBounds(rl.White)
	drawDebugPanel() // bool: showDebugPanel
}

func render() {
	rl.BeginDrawing()
	rl.BeginMode2D(camera)

	drawScene()

	rl.EndMode2D()
	rl.EndDrawing()
}

func initialize() {
	rl.InitWindow(screenWidth, screenHeight, "hello raylib")
	rl.SetTargetFPS(FPS)
	rl.SetExitKey(0)

	cameraTarget.X = float32(screenWidth/2)
	cameraTarget.Y = float32(screenHeight/2)
	cameraOffset.X = 0
	cameraOffset.Y = 0
	camera = rl.NewCamera2D(cameraTarget, cameraOffset, 0, 1)

	perlinImg = rl.GenImagePerlinNoise(800, 800, 0, 0, 5)
	perlinTexture = rl.LoadTextureFromImage(perlinImg)

	
	gopherTexture = rl.LoadTexture("res/panic.png")
	gopher.SetBounds(0, 0, 800, 800, 10)
}

func exit() {
	defer rl.CloseWindow()

	rl.UnloadTexture(perlinTexture)
	rl.UnloadTexture(gopherTexture)
}

func main() {
	initialize()


	for running {
		input()
		update()
		render()
	}
	exit()
}
