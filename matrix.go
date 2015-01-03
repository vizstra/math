package math

import (
	"errors"
	"fmt"
)

const (
	// Column Major Ordering
	// T = Translation/Location
	X1 = iota // 0
	Y1        // 1
	Z1        // 2
	T1        // 3
	X2        // 4
	Y2        // 5
	Z2        // 6
	T2        // 7
	X3        // 8
	Y3        // 9
	Z3        // 10
	T3        // 11
	X4        // 12
	Y4        // 13
	Z4        // 14
	T4        // 15
)

type Mat struct {
	X1, Y1, Z1, T1 float64
	X2, Y2, Z2, T2 float64
	X3, Y3, Z3, T3 float64
	X4, Y4, Z4, T4 float64
}

func CreateMatrixFromVec(a, b, c Vec) Mat {
	return Mat{a.X, a.Y, a.Z, 0,
		b.X, b.Y, b.Z, 0,
		c.X, c.Y, c.Z, 0,
		0, 0, 0, 0}
}

// String returns the string representation of the Mat.
//
// [Tested]
// 20JUN12
//
// [Performance]
// 20JUN12 : 66 - 103 microseconds : Mercury
func (m *Mat) String() string {
	s := fmt.Sprintf("| %-10f %-10f %-10f %-10f |\n", m.X1, m.X2, m.X3, m.X4)
	s += fmt.Sprintf("| %-10f %-10f %-10f %-10f |\n", m.Y1, m.Y2, m.Y3, m.Y4)
	s += fmt.Sprintf("| %-10f %-10f %-10f %-10f |\n", m.Z1, m.Z2, m.Z3, m.Z4)
	s += fmt.Sprintf("| %-10f %-10f %-10f %-10f |\n", m.T1, m.T2, m.T3, m.T4)
	return s
}

// Identity returns a 4x4 identity matrix.
//
// [Tested]
// 20JUN12
//
// [Performance]
// 20JUN12 : 28 nsec/call : Mercury
//
// (see Mat.Mutiply for considerations)
// 1 million iterations returning ref 0.095409
// 1 million iterations returning val 0.027894
func Identity() Mat {
	return Mat{1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1}
}

// Multiply multiply two Mat matrices and
// returns the resultant Mat.
//
// [Tested]
// 20JUN12
//
// [Performance]
// 20JUN12 : 72 nsec/call : Mercury
//
// Notes: When performance testing was being conducted
//        we looked at passing refs vs values
//
//        func (a *Mat) Mulitply(b *Mat) *Mat
//				vs.
//        func (a Mat) Mulitply(b Mat) Mat
//
//        1 million iterations by ref was 0.117456 secs
//        1 million iterations by value was 0.07239 secs
//
func (a Mat) Multiply(b Mat) Mat {

	m := Mat{}

	m.X1 = a.X1*b.X1 + a.X2*b.Y1 + a.X3*b.Z1 + a.X4*b.T1
	m.Y1 = a.Y1*b.X1 + a.Y2*b.Y1 + a.Y3*b.Z1 + a.Y4*b.T1
	m.Z1 = a.Z1*b.X1 + a.Z2*b.Y1 + a.Z3*b.Z1 + a.Z4*b.T1
	m.T1 = a.T1*b.X1 + a.T2*b.Y1 + a.T3*b.Z1 + a.T4*b.T1

	m.X2 = a.X1*b.X2 + a.X2*b.Y2 + a.X3*b.Z2 + a.X4*b.T2
	m.Y2 = a.Y1*b.X2 + a.Y2*b.Y2 + a.Y3*b.Z2 + a.Y4*b.T2
	m.Z2 = a.Z1*b.X2 + a.Z2*b.Y2 + a.Z3*b.Z2 + a.Z4*b.T2
	m.T2 = a.T1*b.X2 + a.T2*b.Y2 + a.T3*b.Z2 + a.T4*b.T2

	m.X3 = a.X1*b.X3 + a.X2*b.Y3 + a.X3*b.Z3 + a.X4*b.T3
	m.Y3 = a.Y1*b.X3 + a.Y2*b.Y3 + a.Y3*b.Z3 + a.Y4*b.T3
	m.Z3 = a.Z1*b.X3 + a.Z2*b.Y3 + a.Z3*b.Z3 + a.Z4*b.T3
	m.T3 = a.T1*b.X3 + a.T2*b.Y3 + a.T3*b.Z3 + a.T4*b.T3

	m.X4 = a.X1*b.X4 + a.X2*b.Y4 + a.X3*b.Z4 + a.X4*b.T4
	m.Y4 = a.Y1*b.X4 + a.Y2*b.Y4 + a.Y3*b.Z4 + a.Y4*b.T4
	m.Z4 = a.Z1*b.X4 + a.Z2*b.Y4 + a.Z3*b.Z4 + a.Z4*b.T4
	m.T4 = a.T1*b.X4 + a.T2*b.Y4 + a.T3*b.Z4 + a.T4*b.T4

	return m
}

// TODO This need to be tested
func (a Mat) Translate(b Vec) Mat {
	m := a
	m.X4 = a.X1*b.X + a.X2*b.Y + a.X3*b.Z + a.X4*b.W
	m.Y4 = a.Y1*b.X + a.Y2*b.Y + a.Y3*b.Z + a.Y4*b.W
	m.Z4 = a.Z1*b.X + a.Z2*b.Y + a.Z3*b.Z + a.Z4*b.W
	m.T4 = 1
	return m
}

// CalculateNormalMatrix will attempt to calculate a
// normal matrix, but if the determinant ends up
// being 0, it will return an error.
func (a Mat) CalculateNormalMatrix() (Mat3, error) {
	m := a.Mat3()

	m, err := m.Inverse()
	if err != nil {
		return m, err
	}

	return m.Transpose(), nil
}

// Array return a 2D representation of the Mat
func (a Mat) Array() [4][4]float64 {
	return [4][4]float64{{a.X1, a.Y1, a.Z1, a.T1},
		{a.X2, a.Y2, a.Z2, a.T2},
		{a.X3, a.Y3, a.Z3, a.T3},
		{a.X4, a.Y4, a.Z4, a.T4}}
}

// Given a Mat (4x4), return the determinant of the
// upper-left 3x3 portion of the Mat.
func (a Mat) Mat3() Mat3 {

	return Mat3{
		a.X1, a.Y1, a.Z1,
		a.X2, a.Y2, a.Z2,
		a.X3, a.Y3, a.Z3}
}

// Mat3 is used for special cases like
// normal matrices.
type Mat3 struct {
	X1, Y1, Z1 float64
	X2, Y2, Z2 float64
	X3, Y3, Z3 float64
}

// Det returns the determinant of the Mat3.
func (a Mat3) Det() float64 {

	// aei + bfg + cdh - ceg - bdi - afg
	return a.X1*a.Y2*a.Z3 +
		a.X2*a.Y3*a.Z1 +
		a.X3*a.Y1*a.Z2 -
		a.X3*a.Y2*a.Z1 -
		a.X2*a.Y1*a.Z3 -
		a.X1*a.Y3*a.Z2
}

// Given a Mat (4x4), return the transpose of the
// upper-left 3x3 portion of the Mat as a 3x3 Matrix.
// This is used to support nomral matrix creation.
func (a Mat3) Transpose() Mat3 {

	return Mat3{
		a.X1, a.X2, a.X3,
		a.Y1, a.Y2, a.Y3,
		a.Z1, a.Z2, a.Z3}
}

// Inverse returns the inverse matrix of the Mat3
func (m Mat3) Inverse() (Mat3, error) {
	det := m.Det()

	if det == 0 {
		return m, errors.New("Cannot inverse a matrix with a determinant equal to 0.")
	}

	return m.Adjoint(1 / det), nil
}

//
func (m Mat3) Adjoint(scale float64) Mat3 {

	var b Mat3
	b.X1 = scale * (m.Y2*m.Z2 - m.Y2*m.Z2)
	b.Y1 = scale * (m.Y2*m.Z1 - m.Y1*m.Z2)
	b.Z1 = scale * (m.Y1*m.Z2 - m.Y2*m.Z1)

	b.X2 = scale * (m.X2*m.Z2 - m.X2*m.Z2)
	b.Y2 = scale * (m.X1*m.Z2 - m.X2*m.Z1)
	b.Z2 = scale * (m.X2*m.Z1 - m.X1*m.Z2)

	b.X2 = scale * (m.X2*m.Y2 - m.X2*m.Y2)
	b.Y2 = scale * (m.X2*m.Y1 - m.X1*m.Y2)
	b.Z2 = scale * (m.X1*m.Y2 - m.X2*m.Y1)

	return b
}
