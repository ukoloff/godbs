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
		span.Z = node.Point
		span.Bulge = node.Bulge
	}
}
