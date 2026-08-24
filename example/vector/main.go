// Using the raymath vector functions to bounce a ball around the window.
//
// Usage:
//
//	./build/vector
package main

import (
	"solod.dev/raylib/libraylib"
	"solod.dev/so/c"
)

var _ c.Int // for implicit int -> c.Int conversion

const (
	screenWidth  = 800
	screenHeight = 450
	radius       = 30
)

func main() {
	libraylib.InitWindow(screenWidth, screenHeight, "raylib - vector math")
	defer libraylib.CloseWindow()

	pos := libraylib.Vector2{X: screenWidth / 2, Y: screenHeight / 2}
	speed := libraylib.Vector2{X: 320, Y: 240}

	libraylib.SetTargetFPS(60)

	// Loop until the user closes the window.
	for !libraylib.WindowShouldClose() {
		// Move the ball by speed * elapsed time.
		dt := libraylib.GetFrameTime()
		pos = libraylib.Vector2Add(pos, libraylib.Vector2Scale(speed, dt))

		// Reflect off the walls the ball has reached.
		if pos.X-radius <= 0 || pos.X+radius >= screenWidth {
			speed = libraylib.Vector2Reflect(speed, libraylib.Vector2{X: 1, Y: 0})
		}
		if pos.Y-radius <= 0 || pos.Y+radius >= screenHeight {
			speed = libraylib.Vector2Reflect(speed, libraylib.Vector2{X: 0, Y: 1})
		}

		libraylib.BeginDrawing()
		libraylib.ClearBackground(libraylib.RAYWHITE)
		libraylib.DrawCircleV(pos, radius, libraylib.MAROON)
		libraylib.DrawText("Vector2Add, Vector2Scale and Vector2Reflect from raymath.h", 20, 20, 20, libraylib.DARKGRAY)
		libraylib.EndDrawing()
	}
}
