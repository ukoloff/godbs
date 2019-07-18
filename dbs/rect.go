package dbs

import "math"

// Rect is rectangular shape (size)
type Rect struct {
	Min, Max Point
}

// IsEmpty checks if Rect is empty
func (me *Rect) IsEmpty() bool {
	return math.IsNaN(me.Min.X)
}

// SetEmpty makes Rect empty
func (me *Rect) SetEmpty() {
	me.Min.X = math.NaN()
}
