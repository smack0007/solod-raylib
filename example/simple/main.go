// Using the raylib library to open a window and draw text in it.
//
// Usage:
//
//	make example name=simple
//	./build/simple
package main

import (
	"solod.dev/raylib/libraylib"
	"solod.dev/so/c"
)

var _ c.Int // for implicit int -> c.Int conversion

const (
	screenWidth  = 640
	screenHeight = 320
)

func main() {
	libraylib.InitWindow(screenWidth, screenHeight, "☀️ Solod / Raylib")
	defer libraylib.CloseWindow()

	libraylib.SetTargetFPS(60)

	// Loop until the user closes the window.
	for !libraylib.WindowShouldClose() {
		libraylib.BeginDrawing()
		libraylib.ClearBackground(libraylib.Color{13, 50, 89, 255})
		libraylib.DrawText("solod", 150, 70, 128, libraylib.Color{242, 203, 73, 255})
		libraylib.DrawText("Go can be a better C", 150, 190, 30, libraylib.WHITE)
		libraylib.EndDrawing()
	}
}
