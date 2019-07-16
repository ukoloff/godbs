package dbs

import (
	"os"
	"testing"
)

func TestSaveCircle(t *testing.T) {
	dbs := NewCircle(99)
	f, _ := os.Create("testdata/.!.dbs")
	defer f.Close()
	dbs.Save(f)
}

func TestSaveRect(t *testing.T) {
	dbs := NewRect(&Point{27, 42})
	f, _ := os.Create("testdata/.r.dbs")
	defer f.Close()
	dbs.Save(f)
}

func TestSaveSquare(t *testing.T) {
	dbs := NewSquare(108)
	f, _ := os.Create("testdata/.q.dbs")
	defer f.Close()
	dbs.Save(f)
}
