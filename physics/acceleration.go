package physics

import "fmt"

//To import from other file that's why Acceleration
func GetAcceleration(initialPoint uint, finalPoint uint, time uint) uint {
	var u uint = initialPoint   //initial point in meter uint
	var v uint = finalPoint     //final point in meter uint
	var t uint = time           //time in seconds uint
	subtraction := v - u        //Subtracting initial point from final point
	division := subtraction / t //Dividing subtraction variable by time
	a := division               //Assigning value to a variable
	fmt.Printf("Acceleration is %v ms^2", a)
	return a //return a
}
