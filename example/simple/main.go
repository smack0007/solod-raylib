// Using the raylib library to open a window and draw text in it.
//
// Usage:
//
//	./build/simple
//
// Source: https://github.com/raysan5/raylib/blob/master/examples/core/core_basic_window.c
package main

import (
	"solod.dev/raylib/libraylib"
	"solod.dev/so/c"
)

var _ c.Int // for implicit int -> c.Int conversion

const (
	screenWidth  = 800
	screenHeight = 450
)

func main() {
	libraylib.InitWindow(screenWidth, screenHeight, "raylib - basic window")
	defer libraylib.CloseWindow()

	libraylib.SetTargetFPS(60)

	// Loop until the user closes the window.
	for !libraylib.WindowShouldClose() {
		libraylib.BeginDrawing()
		libraylib.ClearBackground(libraylib.RAYWHITE)
		libraylib.DrawText("Congrats! You created your first window!", 190, 200, 20, libraylib.LIGHTGRAY)
		libraylib.EndDrawing()
	}
}
