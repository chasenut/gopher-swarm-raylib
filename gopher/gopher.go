package gopher

import rl "github.com/gen2brain/raylib-go/raylib"

type Gopher struct {
	Rect rl.Rectangle
	Vel rl.Vector2
}

var (
	GopherTexture rl.Texture2D = rl.LoadTexture("res/panic.png")
	GopherSrc rl.Rectangle = rl.NewRectangle(0, 0, 240, 240)
	Gophers []Gopher =  []Gopher{}
)

func DrawGophers(gophers []Gopher) {
	for _, g := range gophers {
		rl.DrawTexturePro(GopherTexture, 
			GopherSrc, 
			g.Rect,
			rl.NewVector2(g.Rect.X, g.Rect.Y),
			0,
			rl.White)
	}
}

func CreateGopher(rect rl.Rectangle, vel rl.Vector2) {
	g := Gopher {
		Rect: rect,
		Vel: vel,
	}
	Gophers = append(Gophers, g)
}

func (g Gopher) Update(dt float32) {
	g.Rect.X += g.Vel.X * dt
	g.Rect.Y += g.Vel.Y * dt
}

func ClearMemory() {
	rl.UnloadTexture(GopherTexture)
}
