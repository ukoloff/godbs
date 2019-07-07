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

// Apply transforms Path with Matrix & Shift
func (me *Path) Apply(o2 *O2) Path {
	res := make(Path, len(*me))
	for i, node := range *me {
		res[i] = node.Apply(o2)
	}
	return res
}
