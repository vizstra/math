package math

import (
	"math"
)

type Quat struct {
	W, X, Y, Z float64
}

// NewQuat returns a normalized Quat
// which(in this context) is the
// expression of the rotation in a
// rotation vector.  It is used to
// perform rotation and interpolation
// operations.
func NewQuat(radian float64, xyz Vec) Quat {
	var v Quat
	v.W = Cos(radian / 2)
	v.X = Sin(radian/2) * xyz.X
	v.Y = Sin(radian/2) * xyz.Y
	v.Z = Sin(radian/2) * xyz.Z
	return v.Normalize()
}

func NewQuatFromMat(mat Mat) (quat Quat) {
	m := mat.Array()
	var tr float64
	var s float64
	var q [4]float64

	var i, j, k int
	var nxt [3]int = [3]int{1, 2, 0}
	tr = m[0][0] + m[1][1] + m[2][2]

	// check the diagonal
	if tr > 0.0 {
		s = Sqrt(tr + 1.0)
		quat.W = s / 2.0
		s = 0.5 / s
		quat.X = (m[1][2] - m[2][1]) * s
		quat.Y = (m[2][0] - m[0][2]) * s
		quat.Z = (m[0][1] - m[1][0]) * s
	} else {
		// diagonal is negative
		i = 0
		if m[1][1] > m[0][0] {
			i = 1
		}

		if m[2][2] > m[i][i] {
			i = 2
		}

		j = nxt[i]
		k = nxt[j]
		s = Sqrt((m[i][i] - (m[j][j] + m[k][k])) + 1.0)
		q[i] = s * 0.5

		if s != 0.0 {
			s = 0.5 / s
		}

		q[3] = (m[j][k] - m[k][j]) * s
		q[j] = (m[i][j] + m[j][i]) * s
		q[k] = (m[i][k] + m[k][i]) * s
		quat.X = q[0]
		quat.Y = q[1]
		quat.Z = q[2]
		quat.W = q[3]
	}
	return quat
}

// Normalize returns a Quat in the same
// direction, but of length 1.
func (v Quat) Normalize() Quat {
	l := v.Length()
	if l != 0 {
		v.W /= l
		v.X /= l
		v.Y /= l
		v.Z /= l
	}
	return v
}

// Length returns the length of the Quat
func (v Quat) Length() float64 {
	l := float64(math.Sqrt(float64((v.W * v.W) + (v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z))))
	return l
}

func (v Quat) RotMat() Mat {
	var m Mat
	w := v.W
	x := v.X
	y := v.Y
	z := v.Z

	m.X1 = 1 - 2*y*y - 2*z*z
	m.Y1 = 2*x*y - 2*w*z
	m.Z1 = 2*x*z + 2*w*y

	m.X2 = 2*x*y + 2*w*z
	m.Y2 = 1 - 2*x*x - 2*z*z
	m.Z2 = 2*y*z - 2*w*x

	m.X3 = 2*x*z - 2*w*y
	m.Y3 = 2*y*z + 2*w*x
	m.Z3 = 1 - 2*x*x - 2*y*y

	m.T4 = 1.0

	return m
}

func (q Quat) Multiply(r Quat) Quat {
	w := q.W*r.W - q.X*r.X - q.Y*r.Y - q.Z*r.Z
	x := q.W*r.X + q.X*r.W - q.Y*r.Z + q.Z*r.Y
	y := q.W*r.Y + q.X*r.Z + q.Y*r.W - q.Z*r.X
	z := q.W*r.Z - q.X*r.Y + q.Y*r.X + q.Z*r.W
	return Quat{w, x, y, z}
}

func (q Quat) Conjugate() Quat {
	return Quat{q.W, -q.X, -q.Y, -q.Z}
}

func (q Quat) DotProduct(r Quat) float64 {
	return q.W*r.W + q.X*r.X + q.Y*r.Y + q.Z*r.Z
}

func (a Quat) Add(b Quat) Quat {
	return Quat{a.W + b.W, a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func (a Quat) Subtract(b Quat) Quat {
	return Quat{a.W - b.W, a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

// Scales the Quat by the given scalar in all
// dimensions (w, x, y, z)
func (q Quat) Scale(scalar float64) Quat {

	return Quat{q.W * scalar,
		q.X * scalar,
		q.Y * scalar,
		q.Z * scalar}
}

func (q Quat) RotateVector(v Vec) Vec {
	v = v.Normalize()
	vecQuat := Quat{0, v.X, v.Y, v.Z}
	resQuat := vecQuat.Multiply(q.Conjugate())
	resQuat = q.Multiply(resQuat)
	return Vec{0, resQuat.X, resQuat.Y, resQuat.Y}
}

func (q Quat) Rotate(b Quat) Quat {
	//	b        = b.Normalize()
	//	vecQuat := Quat{0, b.X, b.Y, b.Z}
	//	resQuat := vecQuat.Multiply(q.Conjugate())
	//	resQuat  = q.Multiply(resQuat)
	res := Quat{0, q.X * b.X, q.Y * b.Y, q.Z * b.Z}
	return res.Normalize()
}

// Slerp interpolates between two points
// by traveling along the shortest path
// on a sphere with constant velocity.
// Please note that it is non-communicative
// and computationally expensive.
//
func (a Quat) Slerp(b Quat, percent float64) Quat {
	// Normalizing just in case
	a = a.Normalize()
	b = b.Normalize()

	dot := a.DotProduct(b)
	dp := Clamp(dot, -1, 1)
	theta := Acos(dp) * percent
	relq := b.Subtract(a.Scale(dp))
	relq = relq.Normalize()
	relq = a.Scale(Cos(theta)).Add(relq.Scale(Sin(theta)))
	return relq.Normalize()
}

// NLerp handles rotations much more efficiently
// than Slerp.  Does not rotate with constant
// velocity, but that is ok for animations over
// 10Hz.
//
// [Tested]
// 15JUL13
func (a Quat) Nlerp(b Quat, percent float64) Quat {
	a = a.Normalize()
	b = b.Normalize()

	p := 1.0 - percent
	var ret Quat
	if a.DotProduct(b) < 0.0 {
		ret = a.Scale(p).Add(b.Scale(-percent))
	} else {
		ret = a.Scale(p).Add(b.Scale(percent))
	}
	return ret.Normalize()
}
