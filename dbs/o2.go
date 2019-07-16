package dbs

import "math"

// O2 - Transformation matrix
type O2 struct {
	X,
	Y,
	Delta Point
}

// Eye makes transformation matrix identity
func (me *O2) Eye() {
	*me = O2{}
	me.X.X = 1
	me.Y.Y = 1
}

// Det calculates determinant of transforamtion matrix
func (me *O2) Det() float64 {
	return me.X.X*me.Y.Y - me.X.Y*me.Y.X
}

// CCW initialize matrix to rotation
func (me *O2) CCW(angle float64) {
	sin, cos := math.Sincos(angle * math.Pi / 180.0)
	me.X = Point{cos, sin}
	me.Y = Point{-sin, cos}
	me.Delta = Point{}
}
