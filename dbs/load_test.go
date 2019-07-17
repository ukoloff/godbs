package dbs

import (
	"os"
	"testing"
)

func TestLoadCircle(t *testing.T) {
	f, _ := os.Open("testdata/circle.dbs")
	var ci DBS
	ci.Load(f)
}
