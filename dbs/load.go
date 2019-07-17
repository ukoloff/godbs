package dbs

import (
	"bufio"
	"encoding/binary"
	"io"
)

type dbsReader struct {
	reader *bufio.Reader
	unread int
}

// Load DBS from file / Reader
func (me *DBS) Load(from io.Reader) {
	var reader dbsReader
	reader.init(from)
	reader.readRecords()
}

// Constructor
func (me *dbsReader) init(from io.Reader) {
	me.reader = bufio.NewReader(from)
}

// Read binary data
func (me *dbsReader) read(data interface{}) {
	binary.Read(me.reader, binary.LittleEndian, data)
	me.unread -= binary.Size(data)
}

// Read DBS record
func (me *dbsReader) readProlog(data *recAny) bool {
	me.read(&data.recHead)
	if data.Len < 0 {
		return false
	}
	me.read(&data.recTail)
	me.unread = data.payload()
	return true
}

// Consume DBS file
func (me *dbsReader) readRecords() {
	var prolog recAny
	for me.readProlog(&prolog) {
		switch prolog.Type {
		case 1:
			me.rec1(true)
		case 2:
			me.rec1(false)
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
func (me *dbsReader) rec1(isOrig bool) {
}

// Read Record 8
func (me *dbsReader) rec8() {
}

// Read Record 26
func (me *dbsReader) rec26() {
}
