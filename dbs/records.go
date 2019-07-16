package dbs

import (
	"encoding/binary"
)

// Prolog of DBS record
type recHead struct {
	Len int16
}

type recEOF struct {
	recHead
	EOF int16
}

// Rest of general DBS record
type recTail struct {
	ID2,
	Len2,
	_,
	Type,
	_,
	ID,
	_ int16
}

// General DBS record
type recAny struct {
	recHead
	recTail
}

// Path's Node
type rec1item struct {
	X,
	Y,
	Bulge float32
}

type recPoint struct {
	X, Y float32
}

// Transformation matrix
type recO2 struct {
	X, Y, Delta recPoint
}

// Copy of geometry (Path)
type rec2 struct {
	Subtype,
	_,
	Text,
	_,
	AutoSeq,
	_,
	Part,
	_,
	Original,
	_,
	Rev,
	_ int16
	RecO2 recO2
}

// Contours of a Part
type rec8item struct {
	ID,
	_ int16
}

// Name of a Part
type rec26 struct {
	Name [8]byte
}

// Area & perimeter (in decimeters) of a Part
type rec27 struct {
	Area,
	Perimeter float32
}

// Free text about a Part
type rec28 struct {
	Comment [0]byte
}

// Is this End of DBS File?
func (me *recHead) IsEOF() bool {
	return me.Len < 0
}

// Full length of DBS record in bytes
func (me *recHead) bytes() uint32 {
	return uint32(me.Len+1) * 4
}

// End of DBS file
func (me *recEOF) init() {
	me.Len = -1
	me.EOF = -1
}

func (me *rec1item) fromNode(node *Node) {
	me.X = float32(node.X)
	me.Y = float32(node.Y)
	me.Bulge = float32(node.Bulge)
}

func (me *rec1item) Node() Node {
	return Node{
		Point{
			float64(me.X),
			float64(me.Y),
		},
		float64(me.Bulge),
	}
}

func (me *recPoint) Point() Point {
	return Point{
		X: float64(me.X),
		Y: float64(me.Y),
	}
}

func (me *recO2) eye() {
	*me = recO2{}
	me.X.X = 1
	me.Y.Y = 1
}

func (me *recO2) O2() O2 {
	return O2{
		X:     me.X.Point(),
		Y:     me.Y.Point(),
		Delta: me.Delta.Point(),
	}
}

// Payload length for DBS record, bytes
func (me *recAny) payload() uint32 {
	return me.recHead.bytes() - uint32(binary.Size(*me))
}

// Prepare to write
func (me *recAny) beforeWrite(payload int) {
	me.ID2 = me.ID
	me.Len = int16((payload + binary.Size(recAny{}) + 3) / 4)
	me.Len2 = me.Len
}
