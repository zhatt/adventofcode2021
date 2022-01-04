// Package coord implements a coordinate abstraction.
package coord

import (
	"math/big"
)

// Coord data type.
type Coord struct {
	X int
	Y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Add returns a Coord which is the sum of coords.
func Add(coords ...Coord) Coord {
	xsum := 0
	ysum := 0
	for _, coord := range coords {
		xsum += coord.X
		ysum += coord.Y
	}
	return Coord{xsum, ysum}
}

// DistanceManhattan returns the distance between coords.
func DistanceManhattan(coord1, coord2 Coord) int {
	return abs(coord2.X-coord1.X) + abs(coord2.Y-coord1.Y)
}

// MinBound calculates the minimum bounding coordinate of coords.
func MinBound(coords []Coord) Coord {
	if len(coords) == 0 {
		panic("no coords")
	}
	xmin := coords[0].X
	ymin := coords[0].Y
	for _, coord := range coords {
		xmin = min(xmin, coord.X)
		ymin = min(ymin, coord.Y)
	}
	return Coord{xmin, ymin}
}

// MaxBound calculates the maximum bounding coordinate of coords.
func MaxBound(coords []Coord) Coord {
	if len(coords) == 0 {
		panic("no coords")
	}

	xmax := coords[0].X
	ymax := coords[0].Y
	for _, coord := range coords {
		xmax = max(xmax, coord.X)
		ymax = max(ymax, coord.Y)
	}
	return Coord{xmax, ymax}
}

// Slope of line.
type Slope struct {
	Rise int
	Run  int
}

// Normalize a Slope reduce fraction and move negative number to Rise.
func (slope *Slope) Normalize() {
	if slope.Run == 0 {
		// Infinite slope
		slope.Rise = 0
		return
	}
	rat := big.NewRat(int64(slope.Rise), int64(slope.Run))
	slope.Rise = int(rat.Num().Int64())
	slope.Run = int(rat.Denom().Int64())
}

// SlopeNegy calculates the slope for coordinates where positive y goes down.
func SlopeNegy(coord1, coord2 Coord) Slope {
	if coord1.X == coord2.X {
		return Slope{1, 0}
	}

	// Use rational number to normalize slope.
	rat := big.NewRat(int64(-(coord2.Y - coord1.Y)), int64(coord2.X-coord1.X))
	return Slope{int(rat.Num().Int64()), int(rat.Denom().Int64())}
}

// RelationNegy calculate relationship for coordinates where positive y goes down.
//
// Where is coord2 in relation to coord1?
// Returns
// 0:  Same location
// 1:  coord2 is right of coord1
//     coord2 is directly above coord1
// -1:  coord2 is left of coord1.
//      coord2 is directly below coord1
//
func RelationNegy(coord1, coord2 Coord) int {
	if coord1 == coord2 {
		return 0
	}

	if coord1.X != coord2.X {
		if coord2.X > coord1.X {
			return 1
		}
		return -1
	}

	// coord2 is directly above or below coord1.
	if coord2.Y < coord1.Y {
		// Above is considered part of the to the right set.
		return 1
	}

	// Below is considered part of the to the left set.
	return -1
}

func OutOfBounds(coord Coord, bounds []Coord) bool {
	min := MinBound(bounds)
	max := MaxBound(bounds)
	switch {
	case coord.X < min.X:
		return true
	case coord.Y < min.Y:
		return true
	case coord.X > max.X:
		return true
	case coord.Y > max.Y:
		return true
	}
	return false
}
