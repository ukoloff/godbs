package dbs

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	yaml "gopkg.in/yaml.v2"
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

func TestMassLoad(t *testing.T) {
	var data map[string]struct {
		Bounds [2][2]float64 //`yaml:bounds`
		Parts  []struct {
			Name   string
			Bounds [2][2]float64
			Area,
			Perimeter float64
			Paths []struct {
				Closed bool
				Count  int
				Bounds [2][2]float64
				Area,
				Perimeter float64
			}
		}
	}
	info, _ := ioutil.ReadFile("testdata/dbs.yml")
	yaml.Unmarshal(info, &data)
}
