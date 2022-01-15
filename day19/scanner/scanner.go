package scanner

import (
	"fmt"
	"zhatt/aoc2021/aoc"
	"zhatt/aoc2021/coord"
)

type Rotation struct {
	x, y, z int
}

func NewRotation(x, y, z int) Rotation {
	if !(x >= 0 && x <= 3 && y >= 0 && y <= 3 && z >= 0 && z <= 3) {
		panic("bad rotation")
	}

	return Rotation{x: x, y: y, z: z}
}

// These are all possible rotations of the scanner
var Orientations = [24]Rotation{
	// (to Z)
	{0, 0, 0}, // +x
	{0, 0, 1}, // +y
	{0, 0, 2}, // -x
	{0, 0, 3}, // -y

	// (to Z neg)
	{2, 0, 0}, // +x
	{2, 0, 1}, // +y
	{2, 0, 2}, // -x
	{2, 0, 3}, // -y

	// to Y
	{1, 0, 0}, // +x
	{1, 1, 0}, // +z
	{1, 2, 0}, // -x
	{1, 3, 0}, // -z

	// to Y neg
	{3, 0, 0}, // +x
	{3, 1, 0}, // +z
	{3, 2, 0}, // -x
	{3, 3, 0}, // -z

	// to X
	{0, 3, 0}, // -z
	{3, 0, 1}, // +y
	{2, 1, 0}, // +z
	{1, 0, 3}, // -y

	// to X neg
	{0, 1, 0}, // +z
	{1, 0, 1}, // +y
	{0, 3, 2}, // -z
	{3, 0, 3}, // -y
}

type Scanner struct {
	number   int
	location coord.Coord3d // beacon location are relative to this position
	beacons  []coord.Coord3d
	rotation Rotation
}

func New() Scanner {
	return Scanner{}
}

func (s *Scanner) Number() int             { return s.number }
func (s *Scanner) Location() coord.Coord3d { return s.location }

func (s *Scanner) SetRotation(rotation Rotation)      { s.rotation = rotation }
func (s *Scanner) SetLocation(location coord.Coord3d) { s.location = location }

func (s *Scanner) Beacons() []coord.Coord3d {
	b := make([]coord.Coord3d, len(s.beacons))
	copy(b, s.beacons)
	for index, beacon := range b {
		for count := 0; count < s.rotation.x; count++ {
			beacon.Y, beacon.Z = -beacon.Z, beacon.Y
		}
		for count := 0; count < s.rotation.y; count++ {
			beacon.X, beacon.Z = beacon.Z, -beacon.X
		}
		for count := 0; count < s.rotation.z; count++ {
			beacon.X, beacon.Y = -beacon.Y, beacon.X
		}
		beacon.X += s.location.X
		beacon.Y += s.location.Y
		beacon.Z += s.location.Z
		b[index] = beacon
	}
	return b
}

func (s *Scanner) Bounds() (coord.Coord3d, coord.Coord3d) {
	beacons := s.Beacons()
	min := beacons[0]
	max := beacons[0]

	for _, beacon := range beacons {
		if beacon.X < min.X {
			min.X = beacon.X
		}
		if beacon.X > max.X {
			max.X = beacon.X
		}
		if beacon.Y < min.Y {
			min.Y = beacon.Y
		}
		if beacon.Y > max.Y {
			max.Y = beacon.Y
		}
		if beacon.Z < min.Z {
			min.Z = beacon.Z
		}
		if beacon.Z > max.Z {
			max.Z = beacon.Z
		}
	}

	return min, max
}

func ParseInput(inputLines []string) []Scanner {
	inputLines = append(inputLines, "")

	scanners := make([]Scanner, 0)

	startIndex := 0
	for index := range inputLines {
		if inputLines[index] == "" {
			scanner := New()
			scanner.Load(inputLines[startIndex:index])
			startIndex = index + 1

			scanners = append(scanners, scanner)
		}
	}

	return scanners
}

func (s *Scanner) Load(data []string) {

	// Format: --- scanner 0 ---
	_, err := fmt.Sscanf(data[0], "--- scanner %d ---", &s.number)
	aoc.PanicOnError(err)

	for _, line := range data[1:] {
		coord := coord.Coord3d{}

		// Format: 568,-2007,-577
		_, err := fmt.Sscanf(line, "%d,%d,%d",
			&coord.X,
			&coord.Y,
			&coord.Z,
		)
		aoc.PanicOnError(err)

		s.beacons = append(s.beacons, coord)
	}
}

// See if scanner b can be correlated to scanner a by rotating and shifting it.
func Correlate(a, b Scanner) (bool, coord.Coord3d, Rotation) {
	for _, rotation := range Orientations {
		b.SetRotation(rotation)
		correlated, location := correlateLocation(a, b)
		if correlated {
			return true, location, rotation
		}
	}
	return false, coord.Coord3d{}, Rotation{}
}

// See if scanner b can be correlated to scanner a by shifting it.
func correlateLocation(a, b Scanner) (bool, coord.Coord3d) {
	aMinBound, aMaxBound := a.Bounds()

	// NB.  b is a copy of the original Scanner so we don't need to create a
	// copy because we are not modifiying the beacon slice.
	b.SetLocation(coord.Coord3d{X: 0, Y: 0, Z: 0})
	bMinBound, bMaxBound := b.Bounds()

	// Find the location to start and end checking b's beacons correlation
	// with a's beacons.
	minLocationToCheck := coord.Coord3d{X: -(bMaxBound.X - aMinBound.X), Y: -(bMaxBound.Y - aMinBound.Y), Z: -(bMaxBound.Z - aMinBound.Z)}
	maxLocationToCheck := coord.Coord3d{X: -(bMinBound.X - aMaxBound.X), Y: -(bMinBound.Y - aMaxBound.Y), Z: -(bMinBound.Z - aMaxBound.Z)}

	// Create sets of beacons to check against.  We create a set of full
	// beacons and sets projected to x line and xy plain.
	aBeaconSetY0Z0 := make(map[coord.Coord3d]bool)
	aBeaconSetZ0 := make(map[coord.Coord3d]bool)
	aBeaconSet := make(map[coord.Coord3d]bool)
	for _, coord := range a.Beacons() {
		aBeaconSet[coord] = true
		coord.Z = 0
		aBeaconSetZ0[coord] = true
		coord.Y = 0
		aBeaconSetY0Z0[coord] = true
	}

	for x := minLocationToCheck.X; x <= maxLocationToCheck.X; x++ {
		bLocation := coord.Coord3d{X: x, Y: 0, Z: 0}
		b.SetLocation(bLocation)

		// Check for correlation in only X.  If there is no correlation
		// in x then we can skip checking the other directions.
		xCount := 0

		for _, beacon := range b.Beacons() {
			beacon.Y = 0
			beacon.Z = 0
			if aBeaconSetY0Z0[beacon] {
				xCount++
			}
		}

		if xCount < 12 {
			continue
		}

		for y := minLocationToCheck.Y; y <= maxLocationToCheck.Y; y++ {
			bLocation := coord.Coord3d{X: x, Y: y, Z: 0}
			b.SetLocation(bLocation)

			// Check for correlation in only X and Y.  If there is
			// no correlation then skip other direction.
			yCount := 0
			for _, beacon := range b.Beacons() {
				beacon.Z = 0
				if aBeaconSetZ0[beacon] {
					yCount++
				}
			}

			if yCount < 12 {
				continue
			}

			for z := minLocationToCheck.Z; z <= maxLocationToCheck.Z; z++ {

				bLocation := coord.Coord3d{X: x, Y: y, Z: z}
				b.SetLocation(bLocation)

				count := 0
				for _, beacon := range b.Beacons() {
					if aBeaconSet[beacon] {
						count++
					}
				}

				if count >= 12 {
					return true, bLocation
				}
			}
		}
	}
	return false, coord.Coord3d{}
}
