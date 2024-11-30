package physics

import (
	"fmt"
)

// Force formulae
func GetForce(mass, acceleration float64) uint {
	force := mass * acceleration // f =m x a
	forceUint := uint(force)     //conversion of float64 force into uint
	fmt.Printf("Force is %v \n", forceUint)
	return forceUint
}
