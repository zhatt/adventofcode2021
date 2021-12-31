package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleInput = []string{
	"--- scanner 0 ---",
	"404,-588,-901",
	"528,-643,409",
	"-838,591,734",
	"390,-675,-793",
	"-537,-823,-458",
	"-485,-357,347",
	"-345,-311,381",
	"-661,-816,-575",
	"-876,649,763",
	"-618,-824,-621",
	"553,345,-567",
	"474,580,667",
	"-447,-329,318",
	"-584,868,-557",
	"544,-627,-890",
	"564,392,-477",
	"455,729,728",
	"-892,524,684",
	"-689,845,-530",
	"423,-701,434",
	"7,-33,-71",
	"630,319,-379",
	"443,580,662",
	"-789,900,-551",
	"459,-707,401",
	"",
	"--- scanner 1 ---",
	"686,422,578",
	"605,423,415",
	"515,917,-361",
	"-336,658,858",
	"95,138,22",
	"-476,619,847",
	"-340,-569,-846",
	"567,-361,727",
	"-460,603,-452",
	"669,-402,600",
	"729,430,532",
	"-500,-761,534",
	"-322,571,750",
	"-466,-666,-811",
	"-429,-592,574",
	"-355,545,-477",
	"703,-491,-529",
	"-328,-685,520",
	"413,935,-424",
	"-391,539,-444",
	"586,-435,557",
	"-364,-763,-893",
	"807,-499,-711",
	"755,-354,-619",
	"553,889,-390",
	"",
	"--- scanner 2 ---",
	"649,640,665",
	"682,-795,504",
	"-784,533,-524",
	"-644,584,-595",
	"-588,-843,648",
	"-30,6,44",
	"-674,560,763",
	"500,723,-460",
	"609,671,-379",
	"-555,-800,653",
	"-675,-892,-343",
	"697,-426,-610",
	"578,704,681",
	"493,664,-388",
	"-671,-858,530",
	"-667,343,800",
	"571,-461,-707",
	"-138,-166,112",
	"-889,563,-600",
	"646,-828,498",
	"640,759,510",
	"-630,509,768",
	"-681,-892,-333",
	"673,-379,-804",
	"-742,-814,-386",
	"577,-820,562",
	"",
	"--- scanner 3 ---",
	"-589,542,597",
	"605,-692,669",
	"-500,565,-823",
	"-660,373,557",
	"-458,-679,-417",
	"-488,449,543",
	"-626,468,-788",
	"338,-750,-386",
	"528,-832,-391",
	"562,-778,733",
	"-938,-730,414",
	"543,643,-506",
	"-524,371,-870",
	"407,773,750",
	"-104,29,83",
	"378,-903,-323",
	"-778,-728,485",
	"426,699,580",
	"-438,-605,-362",
	"-469,-447,-387",
	"509,732,623",
	"647,635,-688",
	"-868,-804,481",
	"614,-800,639",
	"595,780,-596",
	"",
	"--- scanner 4 ---",
	"727,592,562",
	"-293,-554,779",
	"441,611,-461",
	"-714,465,-776",
	"-743,427,-804",
	"-660,-479,-426",
	"832,-632,460",
	"927,-485,-438",
	"408,393,-506",
	"466,436,-512",
	"110,16,151",
	"-258,-428,682",
	"-393,719,612",
	"-211,-452,876",
	"808,-476,-593",
	"-575,615,604",
	"-485,667,467",
	"-680,325,-822",
	"-627,-443,-432",
	"872,-547,-609",
	"833,512,582",
	"807,604,487",
	"839,-516,451",
	"891,-625,532",
	"-652,-548,-490",
	"30,-46,-14",
}

func TestDistanceManhattan(t *testing.T) {
	coord1 := Coord{1, 2, 3}
	coord2 := Coord{5, 6, 7}
	assert.Equal(t, 4+4+4, DistanceManhattan(coord1, coord2))
}

func TestParseInput(t *testing.T) {
	scanners := ParseInput(exampleInput)
	assert.Equal(t, 5, len(scanners))

	assert.Equal(t, 4, scanners[4].Number())
	assert.Equal(t, Coord{}, scanners[4].Location())
	b := scanners[4].Beacons()
	assert.Equal(t, Coord{X: 30, Y: -46, Z: -14}, b[len(b)-1])
}

func TestLoad(t *testing.T) {
	data := []string{
		"--- scanner 0 ---",
		"404,-588,-901",
		"528,-643,409",
		"-838,591,734",
	}

	scanner := New()
	scanner.Load(data)
	beacons := scanner.Beacons()

	assert.Equal(t, 3, len(beacons))
	assert.Contains(t, beacons, Coord{X: 404, Y: -588, Z: -901})
	assert.Contains(t, beacons, Coord{X: 528, Y: -643, Z: 409})
	assert.Contains(t, beacons, Coord{X: -838, Y: 591, Z: 734})
	assert.NotContains(t, beacons, Coord{X: 0, Y: 0, Z: 0})
}

func TestRotate(t *testing.T) {
	data := []string{
		"--- scanner 0 ---",
		"1,2,3",
	}

	var tests = []struct {
		x, y, z  int
		expected Coord
	}{
		// No rotation
		{0, 0, 0, Coord{X: 1, Y: 2, Z: 3}},

		// Rotate CCW around x access
		{1, 0, 0, Coord{X: 1, Y: -3, Z: 2}},
		// Verify again to make sure rotation applies to unrotated values
		{1, 0, 0, Coord{X: 1, Y: -3, Z: 2}},
		{2, 0, 0, Coord{X: 1, Y: -2, Z: -3}},
		{3, 0, 0, Coord{X: 1, Y: 3, Z: -2}},

		// Rotate CCW around y access
		{0, 1, 0, Coord{X: 3, Y: 2, Z: -1}},
		// Verify again to make sure rotation applies to unrotated values
		{0, 1, 0, Coord{X: 3, Y: 2, Z: -1}},
		{0, 2, 0, Coord{X: -1, Y: 2, Z: -3}},
		{0, 3, 0, Coord{X: -3, Y: 2, Z: 1}},

		// Rotate CCW around z access
		{0, 0, 1, Coord{X: -2, Y: 1, Z: 3}},
		// Verify again to make sure rotation applies to unrotated values
		{0, 0, 1, Coord{X: -2, Y: 1, Z: 3}},
		{0, 0, 2, Coord{X: -1, Y: -2, Z: 3}},
		{0, 0, 3, Coord{X: 2, Y: -1, Z: 3}},
	}

	scanner := New()
	scanner.Load(data)
	beacons := scanner.Beacons()
	assert.Equal(t, beacons, []Coord{{X: 1, Y: 2, Z: 3}})

	for _, test := range tests {
		scanner.SetRotation(NewRotation(test.x, test.y, test.z))
		beacons = scanner.Beacons()
		assert.Equal(t, 1, len(beacons))
		assert.Equal(t, beacons[0], test.expected)
	}
}

func TestOrientations(t *testing.T) {
	// Verify that the 24 rotations in Orientations give unique beacon
	// positions.
	data := []string{
		"--- scanner 0 ---",
		"1,2,3",
	}

	scanner := New()
	scanner.Load(data)

	coords := make(map[Coord]bool)
	for _, orientation := range Orientations {
		scanner.SetRotation(orientation)
		beacons := scanner.Beacons()
		coords[beacons[0]] = true
	}
	assert.Equal(t, 24, len(coords))
}

func TestCorrelateSelf(t *testing.T) {

	scanners := ParseInput(exampleInput)

	found, location, rotation := Correlate(scanners[0], scanners[0])
	assert.True(t, found)
	assert.Equal(t, Coord{0, 0, 0}, location)
	assert.Equal(t, Rotation{0, 0, 0}, rotation)
}

func TestCorrelateRotation(t *testing.T) {

	scanners := ParseInput(exampleInput)

	found, location, rotation := Correlate(scanners[0], scanners[1])
	assert.True(t, found)
	assert.Equal(t, Coord{68, -1246, -43}, location)
	assert.Equal(t, Rotation{2, 0, 2}, rotation)
}

func TestNotCorrelated(t *testing.T) {

	scanners := ParseInput(exampleInput)

	found, location, rotation := Correlate(scanners[0], scanners[2])
	assert.False(t, found)
	assert.Equal(t, Coord{}, location)
	assert.Equal(t, Rotation{}, rotation)
}
