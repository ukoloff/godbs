package dbs

// O2 - Transformation matrix
type O2 struct {
	X,
	Y,
	Delta Point
}

// Det calculates determinant of transforamtion matrix
func (me *O2) Det() float64 {
	return me.X.X*me.Y.Y - me.X.Y*me.Y.X
}
