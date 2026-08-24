package libraylib_test

import (
	"solod.dev/raylib/libraylib"
)

func ExampleInitWindow() {
	libraylib.InitWindow(800, 450, "raylib - basic window")
	defer libraylib.CloseWindow()

	libraylib.SetTargetFPS(60)
	for !libraylib.WindowShouldClose() {
		libraylib.BeginDrawing()
		libraylib.ClearBackground(libraylib.RAYWHITE)
		libraylib.DrawText("Congrats! You created your first window!", 190, 200, 20, libraylib.LIGHTGRAY)
		libraylib.EndDrawing()
	}
}
