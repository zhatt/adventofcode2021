package coord

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var coord1xyz = Coord3d{1, 3, 5}
var coord2xyz = Coord3d{21, 2, 8}
var coord3xyz = Coord3d{1, 43, 6}

func TestAdd3d(t *testing.T) {
	assert.Equal(t, Coord3d{22, 5, 13}, Add3d(coord1xyz, coord2xyz))
	assert.Equal(t, Coord3d{23, 48, 19}, Add3d(coord1xyz, coord2xyz, coord3xyz))
}

func TestDistanceManhattan3d(t *testing.T) {
	assert.Equal(t, 20+1+3, DistanceManhattan3d(coord1xyz, coord2xyz))
}

func TestMinBound3d(t *testing.T) {
	assert.Equal(t, Coord3d{1, 2, 5}, MinBound3d([]Coord3d{coord1xyz, coord2xyz}))
	assert.Equal(t, Coord3d{1, 2, 5}, MinBound3d([]Coord3d{coord1xyz, coord2xyz, coord3xyz}))
}

func TestMaxBound3d(t *testing.T) {
	assert.Equal(t, Coord3d{1, 43, 6}, MaxBound3d([]Coord3d{coord1xyz, coord3xyz}))
	assert.Equal(t, Coord3d{21, 43, 8}, MaxBound3d([]Coord3d{coord1xyz, coord2xyz, coord3xyz}))
}

func TestOutOfBounds3d(t *testing.T) {
	bounds := []Coord3d{
		{X: 0, Y: 0, Z: 0},
		{X: 3, Y: 2, Z: 6},
		{X: 2, Y: 4, Z: 3},
	}
	assert.False(t, OutOfBounds3d(Coord3d{X: 1, Y: 1, Z: 5}, bounds))
	assert.False(t, OutOfBounds3d(Coord3d{X: 0, Y: 4, Z: 0}, bounds))
	assert.True(t, OutOfBounds3d(Coord3d{X: -1, Y: 4, Z: 10}, bounds))
	assert.True(t, OutOfBounds3d(Coord3d{X: 1, Y: 5, Z: 3}, bounds))
}
