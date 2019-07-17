package dbs

import (
	"math"
	"math/cmplx"
)

// Span of contour - line or arc
type Span struct {
	A     Point
	Bulge float64
	Z     Point
}

func square(x float64) float64 {
	return x * x
}

// Vector returns line direction
func (span *Span) Vector() Point {
	return span.Z.Sub(&span.A)
}

// Area calculates area term for a Span
func (span *Span) Area() float64 {
	s := (span.Z.X*span.A.Y - span.Z.Y*span.A.X) / 2
	if span.Bulge != 0 {
		BuBu := square(span.Bulge)
		s -= (math.Atan(span.Bulge)*square(1+BuBu) - (1-BuBu)*span.Bulge) /
			BuBu / 8 * span.Vector().Abs2()
	}
	return s
}

// Perimeter calculates length of a Span
func (span *Span) Perimeter() float64 {
	p := span.Vector().Abs()
	if span.Bulge != 0 {
		p *= (math.Atan(span.Bulge) / span.Bulge) *
			(1 + square(span.Bulge))
	}
	return p
}

// linear is internal function to find Points in local coordinates
func (span *Span) linear(pos complex128) Point {
	return C2Point(((span.Z.C()-span.A.C())*pos + (span.Z.C() + span.A.C())) / 2)
}

// Zenith is a middle point of arc
func (span *Span) Zenith() Point {
	return span.linear(complex(0, -span.Bulge))
}

// Nadir is a point opposite to a middle point of arc (not on arc itself)
func (span *Span) Nadir() Point {
	return span.linear(complex(0, 1/span.Bulge))
}

// Center returns center of an arc
func (span *Span) Center() Point {
	return span.linear(complex(0, (1/span.Bulge-span.Bulge)/2))
}

// Radius returns that of an arc
func (span *Span) Radius() float64 {
	return math.Abs(1/span.Bulge+span.Bulge) / 4 * span.Vector().Abs()
}

// At finds position of a Point on the Arc
//	-1: Start of Arc
// 	 0: Middle of Arc
//  +1: End of Arc
func (span *Span) At(pos float64) Point {
	return span.linear(
		complex(pos, -span.Bulge) /
			complex(1, -pos*span.Bulge))
}

// AtUniform is like At, returns position of Point on Arc, but slightly more uniform
// (Especially for arcs with big Bulge)
func (span *Span) AtUniform(pos float64) Point {
	q := (math.Sqrt(9+8*square(span.Bulge)) + 1) / 4
	return span.At(pos / (q - (q-1)*square(pos)))
}

// PositionOf finds position for a Point on the Arc
// (Reverse of At)
func (span *Span) PositionOf(point *Point) float64 {
	a := point.Sub(&span.A).Abs()
	z := point.Sub(&span.Z).Abs()
	return (a - z) / (a + z)
}

// tgq - helper (static) function: tan(arg(sqrt(vector)))
func (span *Span) tgq(vector complex128) float64 {
	if real(vector) < 0 {
		return (cmplx.Abs(vector) - real(vector)) / imag(vector)
	}
	return imag(vector) / (cmplx.Abs(vector) + real(vector))
}

// BulgeOf finds bulge for an arc that passes thru a Point
func (span *Span) BulgeOf(point *Point) float64 {
	return span.tgq(
		cmplx.Conj(point.Sub(&span.A).C()) *
			span.Z.Sub(point).C())
}

// LeftBulge finds Bulge for sub-arc from start to that position
func (span *Span) LeftBulge(pos float64) float64 {
	return span.tgq(
		complex(1, span.Bulge) *
			complex(1, pos*span.Bulge))
}

// RightBulge finds Bulge for sub-arc from that position to end
func (span *Span) RightBulge(pos float64) float64 {
	return span.LeftBulge(-pos)
}
