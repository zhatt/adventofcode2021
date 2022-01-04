package coord

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var coord1 = Coord{1, 3}
var coord2 = Coord{21, 2}
var coord3 = Coord{1, 43}

func TestAdd(t *testing.T) {
	assert.Equal(t, Coord{22, 5}, Add(coord1, coord2))
	assert.Equal(t, Coord{23, 48}, Add(coord1, coord2, coord3))
}

func TestDistanceManhattan(t *testing.T) {
	assert.Equal(t, 20+1, DistanceManhattan(coord1, coord2))
}

func TestMinBound(t *testing.T) {
	assert.Equal(t, Coord{1, 2}, MinBound([]Coord{coord1, coord2}))
	assert.Equal(t, Coord{1, 2}, MinBound([]Coord{coord1, coord2, coord3}))
}

func TestMaxBound(t *testing.T) {
	assert.Equal(t, Coord{1, 43}, MaxBound([]Coord{coord1, coord3}))
	assert.Equal(t, Coord{21, 43}, MaxBound([]Coord{coord1, coord2, coord3}))
}

func TestSlope(t *testing.T) {
	expected1 := Slope{2, -1}
	expected1.Normalize()
	assert.Equal(t, expected1, SlopeNegy(Coord{0, 0}, Coord{4, 8}))
	assert.Equal(t, Slope{8, 3}, SlopeNegy(Coord{0, 0}, Coord{3, -8}))
}

func TestRelation(t *testing.T) {
	assert.Equal(t, 0, RelationNegy(Coord{1, 2}, Coord{1, 2}))
	assert.Equal(t, 1, RelationNegy(Coord{1, 2}, Coord{2, 2}))
	assert.Equal(t, -1, RelationNegy(Coord{10, 2}, Coord{2, 2}))
	assert.Equal(t, 1, RelationNegy(Coord{-10, -10}, Coord{-2, -2}))
}

func TestOutOfBounds(t *testing.T) {
	bounds := []Coord{
		{X: 0, Y: 0},
		{X: 3, Y: 2},
		{X: 2, Y: 4},
	}
	assert.False(t, OutOfBounds(Coord{X: 1, Y: 1}, bounds))
	assert.False(t, OutOfBounds(Coord{X: 0, Y: 4}, bounds))
	assert.True(t, OutOfBounds(Coord{X: -1, Y: 4}, bounds))
	assert.True(t, OutOfBounds(Coord{X: 1, Y: 5}, bounds))
}
