package dbs

// Point represents 2D point
type Point struct {
	X, Y float64
}

// C casts Point to complex
func (me *Point) C() complex128 {
	return complex(me.X, me.Y)
}
