package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type kolvo int

const SidesCircle kolvo = 0
const SidesTriangle kolvo = 3
const SidesSquare kolvo = 4

var S float64

func CalcSquare(sideLen float64, sidesNum kolvo) float64 {
	switch sidesNum {
	case SidesCircle:
		S = math.Pi * sideLen * sideLen
	case SidesTriangle:
		S = sideLen * sideLen * math.Sqrt(3) / 4
	case SidesSquare:
		S = sideLen * sideLen
	default:
		S = 0
	}
	return S
}
