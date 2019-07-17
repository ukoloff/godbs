package dbs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadCircle(t *testing.T) {
	f, _ := os.Open("testdata/circle.dbs")
	defer f.Close()
	var ci DBS
	ci.Load(f)
}

func TestLoadAndSave(t *testing.T) {
	files, _ := filepath.Glob("testdata/*.dbs")
	for _, file := range files {
		dir, name := filepath.Split(file)
		if name[0] == '.' {
			continue
		}
		f, _ := os.Open(file)
		defer f.Close()
		var z DBS
		z.Load(f)
		f, _ = os.Create(filepath.Join(dir, "."+name))
		defer f.Close()
		z.Save(f)
	}
}
