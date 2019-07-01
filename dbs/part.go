package dbs

// Part is a named collection of contours
type Part struct {
	Name  string `json:"partid"`
	Paths []Path `json:"paths"`
}

