package dbs

// Path represents polyline made of straight lines and arcs
type Path []Node

// IsClose checks whether Path is closed
func (me *Path) IsClose() bool {
	return len(*me) > 2 && (*me)[0].C() == (*me)[len(*me)-1].C()
}

// Spans enumerates spans on the path
func (me *Path) Spans(callback func(*Span)) {
	var span Span
	for i, node := range *me {
		if 0 != i {
			span.Z = node.Point
			callback(&span)
		}
		span.A = node.Point
		span.Bulge = node.Bulge
	}
}

// Copy makes deep copy
func (me *Path) Copy() Path {
	res := make(Path, len(*me))
	copy(res, *me)
	return res
}

// Reverse returns Path in opposite direction
func (me *Path) Reverse() Path {
	res := make(Path, len(*me))
	for i, node := range *me {
		idx := len(res) - i - 1
		res[idx].Point = node.Point
		if idx == 0 {
			res[i].Bulge = 0
		} else {
			res[idx-1].Bulge = -node.Bulge
		}
	}
	return res
}

// Apply transforms Path with Matrix & Shift
func (me *Path) Apply(o2 *O2) Path {
	res := make(Path, len(*me))
	for i, node := range *me {
		res[i] = node.Apply(o2)
	}
	return res
}

// Area returns area of (closed) Path
func (me *Path) Area() float64 {
	a := 0.0
	if me.IsClose() {
		me.Spans(func(span *Span) {
			a += span.Area()
		})
	}
	return a
}

// Perimeter returns length of a Path
func (me *Path) Perimeter() float64 {
	p := 0.0
	me.Spans(func(span *Span) {
		p += span.Perimeter()
	})
	return p
}
