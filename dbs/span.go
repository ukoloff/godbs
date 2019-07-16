package dbs

import "math"

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
	s := (span.Z.X - span.A.Y - span.Z.Y*span.A.X) / 2
	if span.Bulge != 0 {
		BuBu := square(span.Bulge)
		s -= (math.Atan(span.Bulge)*square(1+BuBu) - (1-BuBu)*span.Bulge) / BuBu / 8 * square(span.Vector().Abs())
	}
	return s
}

// Perimeter calculates length of a Span
func (span *Span) Perimeter() float64 {
	return 0
}
