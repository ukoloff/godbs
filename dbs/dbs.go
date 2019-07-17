package dbs

// DBS is file, containing one or several parts
type DBS []Part

// Copy makes deep copy
func (me *DBS) Copy() DBS {
	res := make(DBS, len(*me))
	for i, part := range *me {
		res[i] = part.Copy()
	}
	return res
}

// Apply transforms DBS file with Matrix & Shift
func (me *DBS) Apply(o2 *O2) DBS {
	res := make(DBS, len(*me))
	for i, part := range *me {
		res[i] = part.Apply(o2)
	}
	return res
}

// Rename changes names of first Parts in file
func (me DBS) Rename(names ...string) DBS {
	for i, name := range names {
		me[i].Rename(name)
	}
	return me
}
