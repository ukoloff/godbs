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

// Size returns dimensions of Rect
func (me *Rect) Size() Point {
	return me.Max.Sub(&me.Min)
}

// AddPoint catches Point into Rect
func (me *Rect) AddPoint(point *Point) {
	if me.IsEmpty() {
		*me = point.Bounds()
	} else {
		if me.Min.X > point.X {
			me.Min.X = point.X
		}
		if me.Min.Y > point.Y {
			me.Min.Y = point.Y
		}
		if me.Max.X < point.X {
			me.Max.X = point.X
		}
		if me.Max.Y < point.Y {
			me.Max.Y = point.Y
		}
	}
}

// AddRect builds union of two Rects
func (me *Rect) AddRect(rect *Rect) {
	me.AddPoint(&rect.Min)
	me.AddPoint(&rect.Max)
}
