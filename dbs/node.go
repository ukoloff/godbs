package dbs

import (
	"encoding/json"
)

// Node is a point on a Path
type Node struct {
	Point
	Bulge float64
}

// UnmarshalJSON implements JSON.parse
func (node *Node) UnmarshalJSON(b []byte) error {
	var points []float64
	if err := json.Unmarshal(b, &points); err != nil {
		return err
	}
	node.X = points[0]
	node.Y = points[1]
	node.Bulge = points[2]
	return nil
}

// MarshalJSON implements JSON.stringify
func (node *Node) MarshalJSON() ([]byte, error) {
	return json.Marshal([]float64{node.X, node.Y, node.Bulge})
}
