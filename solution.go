package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type SidesCount int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

const (
	SidesTriangle = 3
	SidesSquare   = 4
	SidesCircle   = 0
)

func CalcSquare(sideLen float64, sidesNum SidesCount) float64 {
	var area float64

	switch {
	case sidesNum == SidesTriangle:
		area = sideLen * sideLen * math.Sqrt(3) / 4
	case sidesNum == SidesSquare:
		area = sideLen * sideLen
	case sidesNum == SidesCircle:
		area = sideLen * sideLen * math.Pi
	}

	return area
}
