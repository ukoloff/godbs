package dbs

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare(t *testing.T) {
	sq := NewSquare(1000)
	assert.Len(t, sq, 1)
	assert.InDelta(t, sq[0].Area(), 1e6, 1e-3)
	assert.InDelta(t, sq[0].Perimeter(), 4e3, 1e-3)
}

func TestCircle(t *testing.T) {
	ci := NewCircle(10)
	assert.Len(t, ci, 1)
	assert.InDelta(t, ci[0].Area(), math.Pi*1e2, 1e-3)
	assert.InDelta(t, ci[0].Perimeter(), math.Pi*20, 1e-3)
}
