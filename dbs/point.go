package dbs

// Point represents 2D point
type Point struct {
	X, Y float64
}

// C2Point casts complex number to a Point
func C2Point(point complex128) Point {
	return Point{
		X: real(point),
		Y: imag(point),
	}
}

// C casts Point to complex
func (me Point) C() complex128 {
	return complex(me.X, me.Y)
}

// Copy makes (shallow = deep) copy
func (me Point) Copy() Point {
	return me
}

// Expand multiplies Point by number
func (me *Point) Expand(by float64) Point {
	return Point{
		X: me.X * by,
		Y: me.Y * by,
	}
}

// Apply transforms points with Matrix & Shift
func (me *Point) Apply(o2 *O2) Point {
	return C2Point(o2.X.Expand(me.X).C() + o2.Y.Expand(me.Y).C() + o2.Delta.C())
}
