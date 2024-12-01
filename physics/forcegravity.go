package physics

import (
	"fmt"
	"math"
)

// Function to calculate force of gravity
func ForceGravity(mass1 float64, mass2 float64, distance float64) float64 {
	var G float64 = 6.67 * math.Pow(10, -11) //constant value
	fmt.Println(G)
	var force float64 = G * (mass1 * mass2) / (distance * distance) //F= G.m1.m2 / (r)^2
	fmt.Println(force)
	return force
}
