package dbs

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpanBounds(t *testing.T) {
	span := Span{
		A:     Point{0, 1},
		Bulge: 0,
		Z:     Point{1, 0},
	}

	b := span.Bounds()
	assert.InDelta(t, b.Min.Sub(&Point{0, 0}).Abs(), 0, 1e-3)
	assert.InDelta(t, b.Max.Sub(&Point{1, 1}).Abs(), 0, 1e-3)

	span.Bulge = 0.1
	bx := span.Bounds()
	assert.InDelta(t, bx.Min.Sub(&b.Min).Abs(), 0, 1e-3)
	assert.InDelta(t, bx.Max.Sub(&b.Max).Abs(), 0, 1e-3)

	span.Bulge = -0.1
	bx = span.Bounds()
	assert.InDelta(t, bx.Min.Sub(&b.Min).Abs(), 0, 1e-3)
	assert.InDelta(t, bx.Max.Sub(&b.Max).Abs(), 0, 1e-3)

	span.Bulge = 0.42
	bx = span.Bounds()
	assert.InDelta(t, bx.Min.X, bx.Min.X, 1e-3)
	assert.True(t, bx.Min.X < 0)
	assert.InDelta(t, bx.Max.Sub(&b.Max).Abs(), 0, 1e-3)

	span.Bulge = -0.42
	bx = span.Bounds()
	assert.InDelta(t, bx.Min.Sub(&b.Min).Abs(), 0, 1e-3)
	assert.InDelta(t, bx.Max.X, bx.Max.X, 1e-3)
	assert.True(t, bx.Max.X > 1)
}

func TestSpanRadius(t *testing.T) {
	span := Span{
		A:     Point{0, 0},
		Bulge: 1,
		Z:     Point{1, 0},
	}
	assert.InDelta(t, span.Radius(), 0.5, 1e-3)

	span = Span{
		A:     Point{0, 1},
		Bulge: -1,
		Z:     Point{1, 0},
	}
	assert.InDelta(t, span.Radius(), 1/math.Sqrt(2), 1e-3)

	span = Span{
		A:     Point{-1, 1},
		Bulge: 1 / (1 + math.Sqrt(2)),
		Z:     Point{0, 0},
	}
	assert.InDelta(t, span.Radius(), 1, 1e-3)
}

func TestSpanCenter(t *testing.T) {
	span := Span{
		A:     Point{1, 0},
		Bulge: 1,
		Z:     Point{0, 1},
	}
	center := span.Center()
	assert.InDelta(t, center.Sub(&Point{0.5, 0.5}).Abs(), 0, 1e-3)

	span = Span{
		A:     Point{0, 1},
		Bulge: -1 / (1 + math.Sqrt(2)),
		Z:     Point{1, 0},
	}
	center = span.Center()
	assert.InDelta(t, center.Sub(&Point{0, 0}).Abs(), 0, 1e-3)
}

func TestSpanPoints(t *testing.T) {
	span := Span{
		A:     Point{1, 0},
		Bulge: -1,
		Z:     Point{0, 1},
	}
	A := span.At(-1)
	assert.InDelta(t, A.Sub(&span.A).Abs(), 0, 1e-3)
	Z := span.At(+1)
	assert.InDelta(t, Z.Sub(&span.Z).Abs(), 0, 1e-3)

	assert.InDelta(t, span.At(0).Abs(), 0, 1e-3)
	assert.InDelta(t, span.Zenith().Abs(), 0, 1e-3)
	N := span.Nadir()
	assert.InDelta(t, N.Sub(&Point{1, 1}).Abs(), 0, 1e-3)
}

func TestSpanBulge(t *testing.T) {
	N := 10
	span := Span{
		A: Point{1, 0},
		Z: Point{0, 2},
	}

	for span.Bulge = 0; span.Bulge <= 2; span.Bulge++ {
		for j := 0; j <= N; j++ {
			pos := float64(j) / float64(N)
			point := span.At(pos)
			assert.InDelta(t, span.PositionOf(&point), pos, 1e-3)

			bL := span.LeftBulge(pos)
			bR := span.RightBulge(pos)
			assert.InDelta(t, (bL+bR)/(1-bL*bR), span.Bulge, 1e-3)
			assert.True(t, bL >= 0)
			assert.True(t, bR >= 0)
			assert.True(t, bR < 1)

			if j < N {
				assert.InDelta(t, span.BulgeOf(&point), span.Bulge, 1e-3)
			}
		}
	}
}

func TestSpanPerimeter(t *testing.T) {
	span := Span{
		A: Point{0, 0},
		Z: Point{0, 1},
	}
	assert.InDelta(t, span.Perimeter(), 1, 1e-3)
	span.A.Y = 2
	assert.InDelta(t, span.Perimeter(), 1, 1e-3)
	span.A.X = 1
	assert.InDelta(t, span.Perimeter(), math.Sqrt2, 1e-3)
	span.Bulge = math.Sqrt2 - 1
	assert.InDelta(t, span.Perimeter(), math.Pi/2, 1e-3)
	span.Bulge = -span.Bulge
	assert.InDelta(t, span.Perimeter(), math.Pi/2, 1e-3)
	span.Bulge = 1
	assert.InDelta(t, span.Perimeter(), math.Pi/math.Sqrt2, 1e-3)
	span.Bulge = -span.Bulge
	assert.InDelta(t, span.Perimeter(), math.Pi/math.Sqrt2, 1e-3)
}

func TestSpanArea(t *testing.T) {
	span := Span{
		A: Point{0, 0},
		Z: Point{0, 1},
	}
	assert.InDelta(t, span.Area(), 0, 1e-3)
	span.Z.X = 1
	assert.InDelta(t, span.Area(), 0, 1e-3)
	span.A.Y = 1
	assert.InDelta(t, span.Area(), 0.5, 1e-3)

	span = Span{
		A:     Point{10, 0},
		Bulge: 1,
		Z:     Point{11, 0},
	}
	assert.InDelta(t, span.Area(), -math.Pi/8, 1e-3)
	span.Bulge = 1 - math.Sqrt2
	assert.InDelta(t, span.Area(), math.Pi/8-1.0/4, 1e-3)
}

func TestSpanUniform(t *testing.T) {
	N := 10
	span := Span{
		A:     Point{0, 1},
		Bulge: 42,
		Z:     Point{1, 0},
	}
	for ; span.Bulge >= 0; span.Bulge-- {
		var min, max float64
		for j := 0; j <= N; j++ {
			p1 := span.AtUniform(float64(j) / float64(N))
			p2 := span.AtUniform((float64(j) - 0.1) / float64(N))
			dist := p1.Sub(&p2).Abs()
			if j == 0 || min > dist {
				min = dist
			}
			if j == 0 || max < dist {
				max = dist
			}
		}
		assert.True(t, 1-min/max < 0.17)
	}
}
