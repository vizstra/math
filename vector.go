package math

import (
	"fmt"
	. "math"
)

const (
	X = iota
	Y
	Z
	W
)

type Vec struct {
	X, Y, Z, W float64
}

func (v Vec) String() string {
	return fmt.Sprintf("[ %-10f %-10f %-10f %-10f ]", v.X, v.Y, v.Z, v.W)
}

// Scales the Vec by the given (x, y, z)
func (v Vec) Scale(scalar float64) Vec {

	return Vec{v.X * scalar,
		v.Y * scalar,
		v.Z * scalar,
		v.W}
}

// Translates the Vec by the given (x, y, z).
func (v Vec) Translate(x, y, z float64) Vec {
	return Vec{v.X + x,
		v.Y + y,
		v.Z + z,
		v.W}
}

// Length returns the length of the vector.
func (v Vec) Length() float64 {
	return Sqrt((v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z))
}

// Normalize returns a vector in the same
// direction, but of length 1.
func (v Vec) Normalize() Vec {
	l := v.Length()

	if l != 0 {
		v.X /= l
		v.Y /= l
		v.Z /= l
	}

	return v
}

// Subtracts a Vec from this Vec
// and returns the result.
func (a Vec) Subtract(b Vec) Vec {
	return Vec{a.X - b.X,
		a.Y - b.Y,
		a.Z - b.Z,
		a.W}
}

// Adds one vector to another and
// returns the result.
func (a Vec) Add(b Vec) Vec {
	return Vec{a.X + b.X,
		a.Y + b.Y,
		a.Z + b.Z,
		a.W}
}

// CrossProduct returns the vector
// perpendicular to the plane defined
// by the two given vectors.
func (a Vec) CrossProduct(b Vec) Vec {
	return Vec{a.Y*b.Z - a.Z*b.Y,
		-a.X*b.Z + a.Z*b.X,
		a.X*b.Y - a.Y*b.X,
		a.W}
}

//
//
func (a Vec) DotProduct(b Vec) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// Distance returns the distance between two vertices
// in 3D space.
func (a Vec) Distance(b Vec) float64 {
	v := Vec{a.X - b.X, a.Y - b.Y, a.Z - b.Z, a.W}
	return Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

//
//
func CalculateSurfaceNormal(v1, v2, v3 Vec) Vec {
	u := v2.Subtract(v1)
	v := v3.Subtract(v1)
	x := (u.Y * v.Z) - (u.Z * v.Y)
	y := (u.Z * v.X) - (u.X * v.Z)
	z := (u.X * v.Y) - (u.Y * v.X)
	return Vec{x, y, z, 0}
}

// Lerp is the linear interpolation function
func (a Vec) Lerp(b Vec, p float64) Vec {
	return a.Add(b.Subtract(a).Scale(p))
}
