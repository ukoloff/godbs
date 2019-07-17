package dbs

import (
	"bufio"
	"encoding/binary"
	"io"
	"strconv"
)

type dbsReader struct {
	reader    *bufio.Reader
	unread    int
	prolog    recAny
	parts     DBS
	partsMap  map[int16]int
	paths     map[int16][]int16
	copies    map[int16]rec2
	originals map[int16]Path
}

// Load DBS from file / Reader
func (me *DBS) Load(from io.Reader) {
	var reader dbsReader
	reader.init(from)
	reader.readRecords()
	*me = reader.Assemble()
}

// Constructor
func (me *dbsReader) init(from io.Reader) {
	me.reader = bufio.NewReader(from)
	me.partsMap = map[int16]int{}
	me.paths = map[int16][]int16{}
	me.copies = map[int16]rec2{}
	me.originals = map[int16]Path{}
}

// Read binary data
func (me *dbsReader) read(data interface{}) {
	binary.Read(me.reader, binary.LittleEndian, data)
	me.unread -= binary.Size(data)
}

// Read DBS record
func (me *dbsReader) readProlog() bool {
	me.read(&me.prolog.recHead)
	if me.prolog.Len < 0 {
		return false
	}
	me.read(&me.prolog.recTail)
	me.unread = me.prolog.payload()
	return true
}

// Consume DBS file
func (me *dbsReader) readRecords() {
	for me.readProlog() {
		switch me.prolog.Type {
		case 1, 2:
			me.rec1()
		case 8:
			me.rec8()
		case 26:
			me.rec26()
		}
		if me.unread < 0 {
			panic("Error reading DBS")
		} else if me.unread > 0 {
			me.reader.Discard(me.unread)
		}
	}
}

// Read Record 1/2
func (me *dbsReader) rec1() {
	var r2 rec2
	me.read(&r2)
	me.copies[me.prolog.ID] = r2
	if me.prolog.Type != 1 {
		return
	}
	count := me.unread / binary.Size(rec1item{})
	if count < 0 {
		count = 0
	}
	path := make(Path, count)
	me.originals[r2.Original] = path
	for i := range path {
		var node rec1item
		me.read(&node)
		path[i] = node.Node()
	}
}

// Read Record 8
func (me *dbsReader) rec8() {
	me.partByID() // Mark position for Part

	count := me.unread / binary.Size(rec8item{})
	list := make([]int16, count)
	me.paths[me.prolog.ID] = list
	for i := range list {
		var path rec8item
		me.read(&path)
		list[i] = path.ID
	}
}

// Read Record 26
func (me *dbsReader) rec26() {
	var r26 rec26
	me.read(&r26)
	me.partByID().Name = r26.String()
}

// Find or create Part by ID
func (me *dbsReader) partByID() *Part {
	idx, ok := me.partsMap[me.prolog.ID]
	if !ok {
		idx = len(me.parts)
		me.partsMap[me.prolog.ID] = idx
		me.parts = append(me.parts, Part{})
	}
	return &me.parts[idx]
}

// Create (COPY, Rec 2) path by its ID
func (me *dbsReader) pathByID(id int16) Path {
	r2, ok := me.copies[id]
	if !ok {
		panic("DBS Copy not found: " + strconv.Itoa(int(id)))
	}
	orig, ok := me.originals[r2.Original]
	if !ok {
		panic("DBS Original not found: " + strconv.Itoa(int(r2.Original)))
	}
	o2 := r2.RecO2.O2()
	orig = orig.Apply(&o2)
	if r2.Rev != 0 {
		orig = orig.Reverse()
	}
	return orig
}

// Generate DBS
func (me *dbsReader) Assemble() DBS {
	for id, x := range me.paths {
		part := &me.parts[me.partsMap[id]]
		paths := make([]Path, len(x))
		part.Paths = paths
		for i, id := range x {
			paths[i] = me.pathByID(id)
		}
	}
	return me.parts
}
