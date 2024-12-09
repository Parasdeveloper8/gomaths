package mathematics

import (
	"fmt"
	"math"
)

// a + b whole square
func ABWholeSquarePositive(a float64, b float64) float64 {
	wholeSquare := math.Pow(a, 2) + 2*a*b + math.Pow(b, 2)
	fmt.Println(wholeSquare)
	return wholeSquare
}

// a -b whole square
func ABWholeSquareNegative(a float64, b float64) float64 {
	wholeSquare := math.Pow(a, 2) - 2*a*b + math.Pow(b, 2)
	fmt.Println(wholeSquare)
	return wholeSquare
}

// a+b a-b
func AplusBAminusB(a float64, b float64) float64 {
	formula := math.Pow(a, 2) - math.Pow(b, 2)
	fmt.Println(formula)
	return formula
}

// a + b +c whole square
func AplusBplusCwholesquare(a, b, c float64) float64 {
	formula := math.Pow(a, 2) + math.Pow(b, 2) + math.Pow(c, 2) + 2*a*b + 2*b*c + 2*c*a
	fmt.Println(formula)
	return formula
}

// a + b whole cube
func AplusBwholecube(a, b float64) float64 {
	//formula := math.Pow(a, 3) + 3*math.Pow(a, 2)*b + 3*a*math.Pow(b, 2) + math.Pow(b, 3)
	formula := math.Pow(a+b, 3)
	fmt.Println(formula)
	return formula
}
