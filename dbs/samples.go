package dbs

// MakeCircle creates circle of that radius
func (me *DBS) MakeCircle(radius float64) {
	*me = DBS{
		Part{
			Name: "Circle",
			Paths: []Path{
				Path{
					Node{Point{radius, 0}, -1},
					Node{Point{-radius, 0}, -1},
					Node{Point{radius, 0}, 0},
				},
			},
		},
	}
}

// MakeRect creates rectangle of that dimensions
func (me *DBS) MakeRect(size *Point) {
	*me = DBS{
		Part{
			Name: "Rectan",
			Paths: []Path{
				Path{
					Node{Point{0, 0}, 0},
					Node{Point{0, size.Y}, 0},
					Node{Point{size.X, size.Y}, 0},
					Node{Point{size.X, 0}, 0},
					Node{Point{0, 0}, 0},
				},
			},
		},
	}
}

// MakeSquare creates square of that size
func (me *DBS) MakeSquare(size float64) {
	me.MakeRect(&Point{size, size})
	me.Rename("Square")
}
