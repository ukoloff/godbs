package dbs

import "math"

// O2 - Transformation matrix
type O2 struct {
	X,
	Y,
	Delta Point
}

// Eye makes transformation matrix identity
func (me *O2) Eye() *O2 {
	*me = O2{}
	me.X.X = 1
	me.Y.Y = 1
	return me
}

// CCW initialize matrix to rotation
func (me *O2) CCW(angle float64) *O2 {
	sin, cos := math.Sincos(angle * math.Pi / 180.0)
	me.X = Point{cos, sin}
	me.Y = Point{-sin, cos}
	me.Delta = Point{}
	return me
}

// Shift makes pure translation transformation
func (me *O2) Shift(delta Point) *O2 {
	me.Eye()
	me.Delta = delta
	return me
}

// Compose combines several transformation (right to left!)
func (me *O2) Compose(others ...*O2) *O2 {
	var p3 [3]Point
	p3[0].X = 1
	p3[1].Y = 1
	for i := len(others) - 1; i >= 0; i-- {
		t := others[i]
		for j := range p3 {
			p3[j] = p3[j].Apply(t)
		}
	}
	me.X = p3[0].Sub(&p3[2])
	me.Y = p3[1].Sub(&p3[2])
	me.Delta = p3[2]
	return me
}

// Det calculates determinant of transforamtion matrix
func (me *O2) Det() float64 {
	return me.X.X*me.Y.Y - me.X.Y*me.Y.X
}
