package dbs

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSquare(t *testing.T) {
	var sq DBS
	sq.MakeSquare(1000)
	assert.Len(t, sq, 1)
	assert.InDelta(t, sq[0].Area(), 1e6, 1e-3)
	assert.InDelta(t, sq[0].Perimeter(), 4e3, 1e-3)

	testRandomCAE(t, &sq)
}

func TestCircle(t *testing.T) {
	var ci DBS
	ci.MakeCircle(10)
	assert.Len(t, ci, 1)
	assert.InDelta(t, ci[0].Area(), math.Pi*1e2, 1e-3)
	assert.InDelta(t, ci[0].Perimeter(), math.Pi*20, 1e-3)

	testRandomCAE(t, &ci)
}

var rnd *rand.Rand

func testRandomCAE(t *testing.T, dbs *DBS) {
	if rnd == nil {
		rnd = rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	}
	a := (*dbs)[0].Area()
	p := (*dbs)[0].Perimeter()
	for i := 0; i < 100; i++ {
		var part Part
		if i == 0 {
			part = dbs.Copy()[0]
		} else {
			var o2 O2
			o2.CCW(rnd.Float64()*800 - 400)
			o2.Delta.X = rnd.Float64()*200 - 100
			o2.Delta.Y = rnd.Float64()*200 - 100
			part = dbs.Apply(&o2)[0]
		}
		assert.InDelta(t, part.Area(), a, 1e-3)
		assert.InDelta(t, part.Perimeter(), p, 1e-3)
	}

	testReverse(t, dbs)
}

func testReverse(t *testing.T, dbs *DBS) {
	for _, part := range *dbs {
		for _, path := range part.Paths {
			rev := path.Reverse()
			assert.InDelta(t, path.Area(), -rev.Area(), 1e-3)
			assert.InDelta(t, path.Perimeter(), rev.Perimeter(), 1e-3)
		}
	}
}

func TestCircleBounds(t *testing.T) {
	var ci DBS
	ci.MakeCircle(1)
	for angle := 0; angle < 360; angle++ {
		var o2 O2
		o2.CCW(float64(angle))
		z := ci.Apply(&o2)
		b := z.Bounds()
		assert.InDelta(t, b.Min.Sub(&Point{-1, -1}).Abs(), 0, 1e-3)
		assert.InDelta(t, b.Max.Sub(&Point{+1, +1}).Abs(), 0, 1e-3)
	}
}
