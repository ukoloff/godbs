package dbs

// NewCircle creates circle of that radius
func NewCircle(radius float64) DBS {
	return DBS{
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

// NewRect creates rectangle of that dimensions
func NewRect(size *Point) DBS {
	return DBS{
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

// NewSquare creates square of that size
func NewSquare(size float64) DBS {
	return NewRect(&Point{size, size})
}
