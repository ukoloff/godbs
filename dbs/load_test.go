package dbs

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
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

func testBoundsFromArray(bounds *[2][2]float64) Rect {
	return Rect{
		Min: Point{bounds[0][0], bounds[0][0]},
		Max: Point{bounds[1][0], bounds[1][0]},
	}
}

func assertBounds(t *testing.T, a Rect, b [2][2]float64) {
	assert.InDelta(t, b[0][0], a.Min.X, 1e-3)
	assert.InDelta(t, b[0][1], a.Min.Y, 1e-3)
	assert.InDelta(t, b[1][0], a.Max.X, 1e-3)
	assert.InDelta(t, b[1][1], a.Max.Y, 1e-3)
}

func TestMassLoad(t *testing.T) {
	var data map[string]struct {
		Bounds [2][2]float64
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

	for key, info := range data {
		f, _ := os.Open("testdata/" + key + ".dbs")
		defer f.Close()
		var dbs DBS
		dbs.Load(f)
		assertBounds(t, dbs.Bounds(), info.Bounds)

		assert.True(t, len(dbs) == len(info.Parts))
		for i, part := range dbs {
			info := info.Parts[i]
			assert.True(t, info.Name == part.Name)
			assertBounds(t, part.Bounds(), info.Bounds)
			assert.InDelta(t, part.Area(), info.Area, 1e-3)
			assert.InDelta(t, part.Perimeter(), info.Perimeter, 1e-3)

			assert.True(t, len(part.Paths) == len(info.Paths))
			for i, path := range part.Paths {
				info := info.Paths[i]
				assert.True(t, info.Closed == path.IsClose())
				assert.True(t, info.Count == len(path))
				assertBounds(t, path.Bounds(), info.Bounds)
				assert.InDelta(t, path.Area(), info.Area, 1e-3)
				assert.InDelta(t, path.Perimeter(), info.Perimeter, 1e-3)
			}
		}
	}
}
