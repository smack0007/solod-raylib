// Using the raylib library to move a ball with the arrow keys.
//
// Usage:
//
//	./build/input
//
// Source: https://github.com/raysan5/raylib/blob/master/examples/core/core_input_keys.c
package main

import (
	"solod.dev/raylib/libraylib"
	"solod.dev/so/c"
)

var _ c.Int // for implicit int -> c.Int conversion

const (
	screenWidth  = 800
	screenHeight = 450
	ballSpeed    = 2.0
	ballRadius   = 50.0
)

func main() {
	libraylib.InitWindow(screenWidth, screenHeight, "raylib - keyboard input")
	defer libraylib.CloseWindow()

	ball := libraylib.Vector2{X: screenWidth / 2, Y: screenHeight / 2}
	libraylib.SetTargetFPS(60)

	// Loop until the user closes the window.
	for !libraylib.WindowShouldClose() {
		// Move the ball.
		if libraylib.IsKeyDown(libraylib.KEY_RIGHT) {
			ball.X += ballSpeed
		}
		if libraylib.IsKeyDown(libraylib.KEY_LEFT) {
			ball.X -= ballSpeed
		}
		if libraylib.IsKeyDown(libraylib.KEY_UP) {
			ball.Y -= ballSpeed
		}
		if libraylib.IsKeyDown(libraylib.KEY_DOWN) {
			ball.Y += ballSpeed
		}

		// Draw the scene.
		libraylib.BeginDrawing()
		libraylib.ClearBackground(libraylib.RAYWHITE)
		libraylib.DrawText("move the ball with arrow keys", 10, 10, 20, libraylib.DARKGRAY)
		libraylib.DrawCircleV(ball, ballRadius, libraylib.MAROON)
		libraylib.EndDrawing()
	}
}
