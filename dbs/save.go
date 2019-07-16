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

// Write binary data
func (me *dbsWriter) write(data interface{}) {
	binary.Write(me.writer, binary.LittleEndian, data)
}

// Write DBS record
func (me *dbsWriter) startRecord(recType int16, payload int) {
	var rec recAny
	rec.Type = recType
	rec.ID = me.id
	rec.beforeWrite(payload)
	me.write(rec)
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
	var epilog rec2

	me.id++
	me.startRecord(1, binary.Size(epilog)+binary.Size(rec1item{})*len(*path))

	epilog.Subtype = 1
	epilog.Part = partID
	epilog.Original = me.id
	epilog.RecO2.eye()
	me.write(epilog)

	for _, node := range *path {
		var nodeRec rec1item
		nodeRec.fromNode(&node)
		me.write(nodeRec)
	}
}

// Write Rec8 for Part
func (me *dbsWriter) rec8(count int) {
	me.startRecord(8, binary.Size(rec8item{})*count)

	for i := 0; i < count; i++ {
		me.write(rec8item{ID: int16(int(me.id) - count + i)})
	}
}

// Write Rec26 for Part
func (me *dbsWriter) rec26(part *Part) {
	var epilog rec26
	me.startRecord(26, binary.Size(epilog))

	epilog.fromString(part.Name)
	me.write(epilog)
}

// Write Rec27 for Part
func (me *dbsWriter) rec27(part *Part) {
	var epilog rec27
	me.startRecord(27, binary.Size(epilog))

	epilog.Area = float32(part.Area() / 1e4)
	epilog.Perimeter = float32(part.Perimeter() / 1e2)
	me.write(epilog)
}

// Write EOF
func (me *dbsWriter) writeEOF() {
	var eof recEOF
	eof.init()
	me.write(eof)
	me.writer.Flush()
}
