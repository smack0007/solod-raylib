// A flow field. Particles follow a domain-warped value noise and leave
// fading trails on an offscreen canvas.
//
// Usage:
//
//	make example name=flow
//	./build/flow
package main

import (
	"solod.dev/raylib/libraylib"
	"solod.dev/so/c"
	"solod.dev/so/math"
)

var _ c.Int // for implicit int -> c.Int conversion

const (
	width     = 640
	height    = 320
	nParticle = 4000
)

type particle struct {
	pos  libraylib.Vector2
	life float32
}

func main() {
	libraylib.InitWindow(width, height, "☀️ Solod / Raylib")
	defer libraylib.CloseWindow()
	libraylib.SetTargetFPS(60)

	canvas := libraylib.LoadRenderTexture(width, height)
	defer libraylib.UnloadRenderTexture(canvas)
	libraylib.BeginTextureMode(canvas)
	libraylib.ClearBackground(libraylib.Color{8, 8, 14, 255})
	libraylib.EndTextureMode()

	var ps [nParticle]particle
	for i := range nParticle {
		ps[i].pos.X = float32(libraylib.GetRandomValue(0, width))
		ps[i].pos.Y = float32(libraylib.GetRandomValue(0, height))
		ps[i].life = float32(libraylib.GetRandomValue(0, 200))
	}

	t := float32(0)
	for !libraylib.WindowShouldClose() {
		t += 0.003

		libraylib.BeginTextureMode(canvas)
		// Fade the canvas slowly, so the trails decay instead of filling the screen.
		libraylib.DrawRectangle(0, 0, width, height, libraylib.Color{8, 8, 14, 4})
		libraylib.BeginBlendMode(libraylib.BLEND_ADDITIVE)

		for i := range nParticle {
			q := &ps[i]
			x := q.pos.X * 0.004
			y := q.pos.Y * 0.004

			// Domain warp: offset the lookup by another noise sample.
			wx := noise(x+t, y) * 2
			wy := noise(x, y+t) * 2
			ang := float64(noise(x+wx, y+wy) * 6.283 * 2)

			prev := q.pos
			q.pos.X += float32(math.Cos(ang)) * 1.6
			q.pos.Y += float32(math.Sin(ang)) * 1.6

			// Young particles draw dimmer lines.
			s := min(q.life*0.02, 1)
			libraylib.DrawLineV(prev, q.pos, libraylib.Color{
				c.UChar(60 * s), c.UChar(140 * s), c.UChar(200 * s), 60,
			})

			q.life--
			if q.life <= 0 || q.pos.X < 0 || q.pos.X > width ||
				q.pos.Y < 0 || q.pos.Y > height {
				q.pos.X = float32(libraylib.GetRandomValue(0, width))
				q.pos.Y = float32(libraylib.GetRandomValue(0, height))
				q.life = float32(libraylib.GetRandomValue(60, 260))
			}
		}

		libraylib.EndBlendMode()
		libraylib.EndTextureMode()

		libraylib.BeginDrawing()
		tex2D := libraylib.Texture2D{
			Id:      canvas.Texture.Id,
			Width:   canvas.Texture.Width,
			Height:  canvas.Texture.Height,
			Mipmaps: canvas.Texture.Mipmaps,
			Format:  canvas.Texture.Format,
		}
		libraylib.DrawTextureRec(
			tex2D,
			libraylib.Rectangle{0, 0, width, -height},
			libraylib.Vector2{0, 0}, libraylib.WHITE)
		libraylib.EndDrawing()
	}
}

// hash returns a cheap pseudo-random value in [-1, 1) for a grid point.
// The unsigned arithmetic keeps the multiplications from overflowing a
// signed integer.
func hash(x, y int32) float32 {
	n := int32(uint32(x)*374761393 + uint32(y)*668265263)
	n = int32(uint32(n^(n>>13)) * 1274126177)
	return float32((n^(n>>16))&0xffff)/32768 - 1
}

// noise interpolates the four grid hashes around (x, y) with a smoothstep curve.
func noise(x, y float32) float32 {
	xi := int32(math.Floor(float64(x)))
	yi := int32(math.Floor(float64(y)))
	xf := x - float32(xi)
	yf := y - float32(yi)
	u := xf * xf * (3 - 2*xf)
	v := yf * yf * (3 - 2*yf)
	return hash(xi, yi)*(1-u)*(1-v) + hash(xi+1, yi)*u*(1-v) +
		hash(xi, yi+1)*(1-u)*v + hash(xi+1, yi+1)*u*v
}
