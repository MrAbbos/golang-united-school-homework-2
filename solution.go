package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type customeInt int

const (
	SidesCircle   = 0
	SidesTriangle = 3
	SidesSquare   = 4
)

func CalcSquare(sideLen float64, sidesNum customeInt) float64 {
	var result float64
	if sidesNum == SidesTriangle {
		result = (math.Sqrt(3) * sideLen * sideLen) / 4
	} else if sidesNum == SidesSquare {
		result = sideLen * sideLen
	} else if sidesNum == SidesCircle {
		result = math.Pi * sideLen * sideLen
	}
	return result
}
