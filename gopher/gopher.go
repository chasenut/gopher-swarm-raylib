package gopher

import (
	"fmt"

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
	for _, g := range gophers {
		rl.DrawTexturePro(texture, 
			GopherSrc, 
			g.Rect,
			rl.NewVector2(0, 0),
			0,
			rl.White)
	}
}

func AddGopher(gophers *[]Gopher, rect rl.Rectangle, vel rl.Vector2) {
	g := Gopher {
		Rect: rect,
		Vel: vel,
	}
	*gophers = append(*gophers, g)
}

func UpdateGophers(gophers *[]Gopher, dt float32) {
	for _, g := range *gophers {
		fmt.Println("velx: ", g.Rect.X)
		fmt.Println("vely: ", g.Rect.Y)
		g.Rect.X += g.Vel.X * dt * 1000
		g.Rect.Y += g.Vel.Y * dt * 1000
	}
}

