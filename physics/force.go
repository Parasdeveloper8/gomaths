package physics

import (
	"fmt"
)

func GetForce(mass, acceleration float64) uint {
	force := mass * acceleration
	forceUint := uint(force)
	fmt.Printf("Force is %v \n", forceUint)
	return forceUint
}
