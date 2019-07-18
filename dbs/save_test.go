package dbs

import (
	"os"
	"testing"
)

func TestSaveCircle(t *testing.T) {
	var dbs DBS
	dbs.MakeCircle(99)
	f, _ := os.Create("testdata/.c.dbs")
	defer f.Close()
	dbs.Save(f)
}

func TestSaveRect(t *testing.T) {
	var dbs DBS
	dbs.MakeRect(&Point{42, 27})
	f, _ := os.Create("testdata/.r.dbs")
	defer f.Close()
	dbs.Save(f)
}

func TestSaveSquare(t *testing.T) {
	var dbs DBS
	dbs.MakeSquare(108)
	f, _ := os.Create("testdata/.q.dbs")
	defer f.Close()
	dbs.Save(f)
}
