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

	upperBound, lowerBound, rightBound, leftBound rl.Rectangle
)

func SetBounds(x, y, width, height, innerThickness float32) {
	upperBound = rl.NewRectangle(x, y, width, innerThickness)
	lowerBound = rl.NewRectangle(x, y+height-innerThickness, width, innerThickness)
	rightBound = rl.NewRectangle(x+width-innerThickness, y, innerThickness, height)
	leftBound = rl.NewRectangle(x, y, innerThickness, height)
	fmt.Println("new bounds set")
}

func DrawBounds(color rl.Color) {
	rl.DrawRectangle(upperBound.ToInt32().X,
		upperBound.ToInt32().Y,
		upperBound.ToInt32().Width,
		upperBound.ToInt32().Height,
		color)
	rl.DrawRectangle(lowerBound.ToInt32().X,
		lowerBound.ToInt32().Y,
		lowerBound.ToInt32().Width,
		lowerBound.ToInt32().Height,
		color)
	rl.DrawRectangle(rightBound.ToInt32().X,
		rightBound.ToInt32().Y,
		rightBound.ToInt32().Width,
		rightBound.ToInt32().Height,
		color)
	rl.DrawRectangle(leftBound.ToInt32().X,
		leftBound.ToInt32().Y,
		leftBound.ToInt32().Width,
		leftBound.ToInt32().Height,
		color)
}

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
		newX := gophers[i].Rect.X + gophers[i].Vel.X * dt * 10
		newY := gophers[i].Rect.Y + gophers[i].Vel.Y * dt * 10
		newRect := rl.NewRectangle(newX, newY,
			gophers[i].Rect.Width, 
			gophers[i].Rect.Height)
		if rl.CheckCollisionRecs(newRect, upperBound) {
			gophers[i].Vel.Y = -gophers[i].Vel.Y
			newRect.Y = upperBound.Y + upperBound.Height
		} else if rl.CheckCollisionRecs(newRect, lowerBound) {
			gophers[i].Vel.Y = -gophers[i].Vel.Y
			newRect.Y = lowerBound.Y - lowerBound.Height - newRect.Height
		}
		if rl.CheckCollisionRecs(newRect, rightBound) {
			gophers[i].Vel.X = -gophers[i].Vel.X
			newRect.X = rightBound.X - rightBound.Width - newRect.Width
		} else if rl.CheckCollisionRecs(newRect, leftBound) {
			gophers[i].Vel.X = -gophers[i].Vel.X
			newRect.X = leftBound.X + leftBound.Width
		}
		gophers[i].Rect = newRect
	}
}

