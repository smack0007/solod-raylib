// Package libraylib provides low-level Solod bindings for the [raylib] library.
//
// raylib is a simple and easy-to-use library to enjoy videogames programming.
//
// [raylib]: https://www.raylib.com
package libraylib

//so:include <raylib.h>
//so:include <raymath.h>
//so:link raylib

// Angle conversion factors. raylib defines them as float expressions,
// which sobind cannot map, so they are written by hand.
const (
	DEG2RAD = PI / 180
	RAD2DEG = 180 / PI
)
