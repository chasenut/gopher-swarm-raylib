package gopher

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Gopher struct {
	Rect rl.Rectangle
	Vel rl.Vector2
}

var (
	GopherSrc rl.Rectangle = rl.NewRectangle(0, 0, 240, 240)
)

func DrawGophers(gophers []Gopher, texture rl.Texture2D) {
	for i := range gophers {
		rl.DrawTexturePro(texture, 
			GopherSrc, 
			gophers[i].Rect,
			rl.NewVector2(0, 0),
			0,
			rl.White)
	}
}

func AddGopher(gophers []Gopher, rect rl.Rectangle, vel rl.Vector2) []Gopher{
	g := Gopher {
		Rect: rect,
		Vel: vel,
	}
	return append(gophers, g)
}

func UpdateGophers(gophers []Gopher, dt float32) {
	for i := range gophers {
		gophers[i].Rect.X += gophers[i].Vel.X * dt * 10
		gophers[i].Rect.Y += gophers[i].Vel.Y * dt * 10
	}
}

