// Package coord implements a coordinate abstraction.
package coord

// Coord data type.
type Coord3d struct {
	X int
	Y int
	Z int
}

// Add returns a Coord which is the sum of coords.
func Add3d(coords ...Coord3d) Coord3d {
	xsum := 0
	ysum := 0
	zsum := 0
	for _, coord := range coords {
		xsum += coord.X
		ysum += coord.Y
		zsum += coord.Z
	}
	return Coord3d{xsum, ysum, zsum}
}

// DistanceManhattan returns the distance between coords.
func DistanceManhattan3d(coord1, coord2 Coord3d) int {
	return abs(coord2.X-coord1.X) + abs(coord2.Y-coord1.Y) + abs(coord2.Z-coord1.Z)
}

// MinBound calculates the minimum bounding coordinate of coords.
func MinBound3d(coords []Coord3d) Coord3d {
	if len(coords) == 0 {
		panic("no coords")
	}
	xmin := coords[0].X
	ymin := coords[0].Y
	zmin := coords[0].Z
	for _, coord := range coords {
		xmin = min(xmin, coord.X)
		ymin = min(ymin, coord.Y)
		zmin = min(zmin, coord.Z)
	}
	return Coord3d{xmin, ymin, zmin}
}

// MaxBound calculates the maximum bounding coordinate of coords.
func MaxBound3d(coords []Coord3d) Coord3d {
	if len(coords) == 0 {
		panic("no coords")
	}

	xmax := coords[0].X
	ymax := coords[0].Y
	zmax := coords[0].Z
	for _, coord := range coords {
		xmax = max(xmax, coord.X)
		ymax = max(ymax, coord.Y)
		zmax = max(zmax, coord.Z)
	}
	return Coord3d{xmax, ymax, zmax}
}

func OutOfBounds3d(coord Coord3d, bounds []Coord3d) bool {
	min := MinBound3d(bounds)
	max := MaxBound3d(bounds)
	switch {
	case coord.X < min.X:
		return true
	case coord.Y < min.Y:
		return true
	case coord.Z < min.Z:
		return true
	case coord.X > max.X:
		return true
	case coord.Y > max.Y:
		return true
	case coord.Z > max.Z:
		return true
	}
	return false
}
