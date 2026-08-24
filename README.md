# raylib

Solod bindings for [raylib](https://www.raylib.com), a simple and easy-to-use library to enjoy videogames programming.

## Usage

1. Install raylib for your operating system.

2. Install the Solod bindings.

```
go get solod.dev/raylib@latest
```

3. Use it in your code.

```go
package main

import "solod.dev/raylib/libraylib"

func main() {
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
```

## Scope

The bindings cover `raylib.h` (the window, drawing, input, audio and model APIs) and `raymath.h` (vector, matrix and quaternion math). They do not cover `rlgl.h` (the OpenGL abstraction layer).

## Examples

[Basic window](example/simple/main.go)

[Keyboard input](example/input/main.go)

[Vector math](example/vector/main.go)
