package cuboid

import "zhatt/aoc2021/coord"

// Cuboids points are the centers of the point.  We store a set of normalized
// points that are the true vertexes.  I.E. a Cuboid spanning x from 2 to 3
// needs to store 1.5 and 3.5 as the corners.  To above floating point rounding
// error, we will use fixed point and store 10 times the values.
//
// Using the true vertexes simplifies the splitting operations.

type Cuboid struct {
	normalizedMinVertex coord.Coord3d
	normalizedMaxVertex coord.Coord3d
}

func New(minVertex, maxVertex coord.Coord3d) Cuboid {
	normalizedMinVertex := coord.Coord3d{
		X: minVertex.X*10 - 5,
		Y: minVertex.Y*10 - 5,
		Z: minVertex.Z*10 - 5,
	}

	normalizedMaxVertex := coord.Coord3d{
		X: maxVertex.X*10 + 5,
		Y: maxVertex.Y*10 + 5,
		Z: maxVertex.Z*10 + 5,
	}

	return Cuboid{
		normalizedMinVertex: normalizedMinVertex,
		normalizedMaxVertex: normalizedMaxVertex,
	}
}

// Return center of cube vertexes.
func (c *Cuboid) MinVertex() coord.Coord3d {
	return coord.Coord3d{
		X: (c.normalizedMinVertex.X + 5) / 10,
		Y: (c.normalizedMinVertex.Y + 5) / 10,
		Z: (c.normalizedMinVertex.Z + 5) / 10,
	}
}

// Return center of cube vertexes.
func (c *Cuboid) MaxVertex() coord.Coord3d {
	return coord.Coord3d{
		X: (c.normalizedMaxVertex.X - 5) / 10,
		Y: (c.normalizedMaxVertex.Y - 5) / 10,
		Z: (c.normalizedMaxVertex.Z - 5) / 10,
	}
}

// c completely contains other.
func (c Cuboid) Contains(other Cuboid) bool {
	return c.normalizedMinVertex.X <= other.normalizedMinVertex.X &&
		c.normalizedMaxVertex.X >= other.normalizedMaxVertex.X &&
		c.normalizedMinVertex.Y <= other.normalizedMinVertex.Y &&
		c.normalizedMaxVertex.Y >= other.normalizedMaxVertex.Y &&
		c.normalizedMinVertex.Z <= other.normalizedMinVertex.Z &&
		c.normalizedMaxVertex.Z >= other.normalizedMaxVertex.Z
}

func overlapImpl(c1, c2 Cuboid) bool {
	// Check if any corner of c2 is in bounds in c1
	// coord.OutOfBounds3d expects center of cube vertexes.
	for _, x := range []int{c2.MinVertex().X, c2.MaxVertex().X} {
		for _, y := range []int{c2.MinVertex().Y, c2.MaxVertex().Y} {
			for _, z := range []int{c2.MinVertex().Z, c2.MaxVertex().Z} {

				vertex := coord.Coord3d{X: x, Y: y, Z: z}

				if !(coord.OutOfBounds3d(vertex, []coord.Coord3d{c1.MinVertex(), c1.MaxVertex()})) {
					return true
				}
			}
		}
	}

	return false
}

func Overlap(c1, c2 Cuboid) bool {
	return overlapImpl(c1, c2) || overlapImpl(c2, c1)
}

func (c Cuboid) containsNormalizedX(normalizedX int) bool {
	// must be inside and not just on the edge
	return c.normalizedMinVertex.X < normalizedX &&
		c.normalizedMaxVertex.X > normalizedX

}

func (c Cuboid) containsNormalizedY(normalizedY int) bool {
	// must be inside and not just on the edge
	return c.normalizedMinVertex.Y < normalizedY &&
		c.normalizedMaxVertex.Y > normalizedY

}

func (c Cuboid) containsNormalizedZ(normalizedZ int) bool {
	// must be inside and not just on the edge
	return c.normalizedMinVertex.Z < normalizedZ &&
		c.normalizedMaxVertex.Z > normalizedZ

}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (c Cuboid) splitX(normalizedX int) []Cuboid {

	if absInt(normalizedX%10) != 5 {
		panic("bad split - not aligned")
	}

	if !c.containsNormalizedX(normalizedX) {
		panic("bad split - outside of cuboid")
	}

	min := c.normalizedMinVertex
	max := c.normalizedMaxVertex
	max.X = normalizedX

	cuboid1 := Cuboid{normalizedMinVertex: min, normalizedMaxVertex: max}

	min = c.normalizedMinVertex
	max = c.normalizedMaxVertex
	min.X = normalizedX

	cuboid2 := Cuboid{normalizedMinVertex: min, normalizedMaxVertex: max}

	return []Cuboid{cuboid1, cuboid2}
}

func (c Cuboid) splitY(normalizedY int) []Cuboid {

	if absInt(normalizedY%10) != 5 {
		panic("bad split - not aligned")
	}

	if !c.containsNormalizedY(normalizedY) {
		panic("bad split - outside of cuboid")
	}

	min := c.normalizedMinVertex
	max := c.normalizedMaxVertex
	max.Y = normalizedY

	cuboid1 := Cuboid{normalizedMinVertex: min, normalizedMaxVertex: max}

	min = c.normalizedMinVertex
	max = c.normalizedMaxVertex
	min.Y = normalizedY

	cuboid2 := Cuboid{normalizedMinVertex: min, normalizedMaxVertex: max}

	return []Cuboid{cuboid1, cuboid2}
}

func (c Cuboid) splitZ(normalizedZ int) []Cuboid {

	if absInt(normalizedZ%10) != 5 {
		panic("bad split - not aligned")
	}

	if !c.containsNormalizedZ(normalizedZ) {
		panic("bad split - outside of cuboid")
	}

	min := c.normalizedMinVertex
	max := c.normalizedMaxVertex
	max.Z = normalizedZ

	cuboid1 := Cuboid{normalizedMinVertex: min, normalizedMaxVertex: max}

	min = c.normalizedMinVertex
	max = c.normalizedMaxVertex
	min.Z = normalizedZ

	cuboid2 := Cuboid{normalizedMinVertex: min, normalizedMaxVertex: max}

	return []Cuboid{cuboid1, cuboid2}
}

func (c Cuboid) Volume() int {
	return ((c.normalizedMaxVertex.X - c.normalizedMinVertex.X) / 10) *
		((c.normalizedMaxVertex.Y - c.normalizedMinVertex.Y) / 10) *
		((c.normalizedMaxVertex.Z - c.normalizedMinVertex.Z) / 10)
}
