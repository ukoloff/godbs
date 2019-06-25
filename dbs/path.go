package dbs

// Path represents polyline made of straight lines and arcs
type Path []Node

// IsClose checks whether Path is closed
func (me *Path) IsClose() bool {
	return len(*me) > 2 && (*me)[0].C() == (*me)[len(*me)-1].C()
}
