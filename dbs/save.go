package dbs

import (
	"bufio"
	"encoding/binary"
	"io"
)

type dbsWriter struct {
	writer *bufio.Writer
	id     int16
}

// Save DBS to writer stream
func (me *DBS) Save(dst io.Writer) {
	var saver dbsWriter
	saver.init(dst)
	defer saver.writer.Flush()
	saver.writeDBS(me)
}

// Constructor for dbsWriter
func (me *dbsWriter) init(from io.Writer) {
	me.writer = bufio.NewWriter(from)
}

// Write DBS file
func (me *dbsWriter) writeDBS(dbs *DBS) {
	for _, part := range *dbs {
		me.writePart(&part)
	}
	me.writeEOF()
}

// Write a Part
func (me *dbsWriter) writePart(part *Part) {
	partID := me.id + int16(len(part.Paths)) + 1
	for i, path := range part.Paths {
		partID := partID
		if i == 0 {
			partID = -partID
		}
		me.rec1(&path, partID)
	}
	me.id++ // = partID
	me.rec8(len(part.Paths))
	me.rec26(part)
	me.rec27(part)
}

// Write Rec1 for Path
func (me *dbsWriter) rec1(path *Path, partID int16) {
	var prolog recAny
	var epilog rec2

	me.id++
	prolog.ID = me.id
	prolog.Type = 1
	prolog.beforeWrite(binary.Size(epilog) + binary.Size(rec1item{})*len(*path))
	epilog.Subtype = 1
	epilog.Part = partID
	epilog.Original = me.id
	epilog.RecO2.eye()

	binary.Write(me.writer, binary.LittleEndian, prolog)
	binary.Write(me.writer, binary.LittleEndian, epilog)

	for _, node := range *path {
		var nodeRec rec1item
		nodeRec.fromNode(&node)
		binary.Write(me.writer, binary.LittleEndian, nodeRec)
	}
}

// Write Rec8 for Part
func (me *dbsWriter) rec8(count int) {
	var prolog recAny
	prolog.ID = me.id
	prolog.Type = 8
	prolog.beforeWrite(binary.Size(rec8item{}) * count)
	binary.Write(me.writer, binary.LittleEndian, prolog)

	for i := 0; i < count; i++ {
		ref := rec8item{ID: int16(int(me.id) - count + i)}
		binary.Write(me.writer, binary.LittleEndian, ref)
	}
}

// Write Rec26 for Part
func (me *dbsWriter) rec26(part *Part) {
	var prolog recAny
	var epilog rec26
	prolog.ID = me.id
	prolog.Type = 26
	prolog.beforeWrite(binary.Size(epilog))
	binary.Write(me.writer, binary.LittleEndian, prolog)

	epilog.fromString(part.Name)

	binary.Write(me.writer, binary.LittleEndian, epilog)
}

// Write Rec27 for Part
func (me *dbsWriter) rec27(part *Part) {
	var prolog recAny
	var epilog rec27
	prolog.ID = me.id
	prolog.Type = 27
	prolog.beforeWrite(binary.Size(epilog))
	binary.Write(me.writer, binary.LittleEndian, prolog)

	epilog.Area = float32(part.Area() / 1e4)
	epilog.Perimeter = float32(part.Perimeter() / 1e2)

	binary.Write(me.writer, binary.LittleEndian, epilog)
}

// Write EOF
func (me *dbsWriter) writeEOF() {
	var eof recEOF
	eof.init()
	binary.Write(me.writer, binary.LittleEndian, eof)
	defer me.writer.Flush()
}
