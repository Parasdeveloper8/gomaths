package main

import (
	"math"

	"github.com/Parasdeveloper8/gomaths.git/physics"
)

func main() {
	physics.ForceGravity(6*math.Pow(10, 24), 7.4*math.Pow(10, 22), 384*math.Pow(10, 8))

}
