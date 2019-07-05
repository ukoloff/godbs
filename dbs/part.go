package dbs

// Part is a named collection of contours
type Part struct {
	Name  string `json:"partid"`
	Paths []Path `json:"paths"`
}

// Copy makes deep copy
func (me *Part) Copy() Part {
	res := *me
	for i, path := range res.Paths {
		res.Paths[i] = path.Copy()
	}
	return res
}

// Apply transforms Part with Matrix & Shift
func (me *Part) Apply(o2 *O2) Part {
	res := *me
	for i, path := range res.Paths {
		res.Paths[i] = path.Apply(o2)
	}
	return res
}
