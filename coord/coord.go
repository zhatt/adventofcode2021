// Package coord implements a coordinate abstraction.
package coord

import (
	"math/big"
)

// Coord data type.
type Coord struct {
	Xval int
	Yval int
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
		xsum += coord.Xval
		ysum += coord.Yval
	}
	return Coord{xsum, ysum}
}

// DistanceManhattan returns the distance between coords.
func DistanceManhattan(coord1, coord2 Coord) int {
	return abs(coord2.Xval-coord1.Xval) + abs(coord2.Yval-coord1.Yval)
}

// MinBound calculates the minimum bounding coordinate of coords.
func MinBound(coords []Coord) Coord {
	if len(coords) == 0 {
		panic("no coords")
	}
	xmin := coords[0].Xval
	ymin := coords[0].Yval
	for _, coord := range coords {
		xmin = min(xmin, coord.Xval)
		ymin = min(ymin, coord.Yval)
	}
	return Coord{xmin, ymin}
}

// MaxBound calculates the maximum bounding coordinate of coords.
func MaxBound(coords []Coord) Coord {
	if len(coords) == 0 {
		panic("no coords")
	}

	xmax := coords[0].Xval
	ymax := coords[0].Yval
	for _, coord := range coords {
		xmax = max(xmax, coord.Xval)
		ymax = max(ymax, coord.Yval)
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
	if coord1.Xval == coord2.Xval {
		return Slope{1, 0}
	}

	// Use rational number to normalize slope.
	rat := big.NewRat(int64(-(coord2.Yval - coord1.Yval)), int64(coord2.Xval-coord1.Xval))
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

	if coord1.Xval != coord2.Xval {
		if coord2.Xval > coord1.Xval {
			return 1
		}
		return -1
	}

	// coord2 is directly above or below coord1.
	if coord2.Yval < coord1.Yval {
		// Above is considered part of the to the right set.
		return 1
	}

	// Below is considered part of the to the left set.
	return -1
}
