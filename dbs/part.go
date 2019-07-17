package dbs

// Part is a named collection of contours
type Part struct {
	Name  string `json:"partid"`
	Paths []Path `json:"paths"`
}

// dup used internally to duplicate a Part
func (me *Part) dup(dup func(path *Path) Path) Part {
	res := *me
	res.Paths = make([]Path, len(me.Paths))
	copy(res.Paths, me.Paths)
	for i, path := range res.Paths {
		res.Paths[i] = dup(&path)
	}
	return res
}

// Copy makes deep copy
func (me *Part) Copy() Part {
	return me.dup(func(path *Path) Path {
		return path.Copy()
	})
}

// Apply transforms Part with Matrix & Shift
func (me *Part) Apply(o2 *O2) Part {
	return me.dup(func(path *Path) Path {
		return path.Apply(o2)
	})
}

// Area returns area of a Part
func (me *Part) Area() float64 {
	a := 0.0
	for _, path := range me.Paths {
		a += path.Area()
	}
	return a
}

// Perimeter returns length of a Part
func (me *Part) Perimeter() float64 {
	p := 0.0
	for _, path := range me.Paths {
		p += path.Perimeter()
	}
	return p
}

// Rename changes name of Part
func (me *Part) Rename(name string) Part {
	me.Name = name
	return *me
}
