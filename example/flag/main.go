// Using the raylib library to open a window and draw a waving flag.
//
// Usage:
//
//	make example name=flag
//	./build/flag
package main

import (
	rl "solod.dev/raylib/libraylib"
	"solod.dev/so/c"
	"solod.dev/so/math"
)

var _ c.Int // for implicit int -> c.Int conversion

const (
	screenWidth  = 800
	screenHeight = 450

	halfScreenWidth  = screenWidth / 2
	halfScreenHeight = screenHeight / 2

	flagWidth  = 640
	flagHeight = 320

	halfFlagWidth  = flagWidth / 2
	halfFlagHeight = flagHeight / 2

	flagParts     = 25
	flagPartWidth = flagWidth / flagParts

	elapsedTimeFactor = 10
	yOffsetFactor     = 10
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Solod Flag")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	image := rl.GenImageColor(flagWidth, flagHeight, rl.Color{R: 13, G: 50, B: 89, A: 255})
	rl.ImageDrawText(&image, "solod", 150, 70, 128, rl.Color{R: 242, G: 203, B: 73, A: 255})
	rl.ImageDrawText(&image, "Go can be a better C", 150, 190, 30, rl.WHITE)

	texture := rl.LoadTextureFromImage(image)

	rl.UnloadImage(image)

	for !rl.WindowShouldClose() {
		elapsedTime := rl.GetTime()

		rl.BeginDrawing()
		rl.ClearBackground(rl.BLACK)
		for i := 0; i < flagParts; i += 1 {
			angle := float64(flagPartWidth*i) + (elapsedTime * elapsedTimeFactor)
			rl.DrawTexturePro(
				texture,
				rl.Rectangle{
					X:      float32(flagPartWidth * i),
					Y:      0,
					Width:  flagPartWidth,
					Height: flagHeight,
				},
				rl.Rectangle{
					X:      halfScreenWidth + float32(flagPartWidth*i),
					Y:      halfScreenHeight + (float32(math.Sin(angle)) * yOffsetFactor),
					Width:  flagPartWidth,
					Height: flagHeight,
				},
				rl.Vector2{X: halfFlagWidth, Y: halfFlagHeight},
				0.0,
				rl.WHITE)
		}
		rl.EndDrawing()
	}
}
