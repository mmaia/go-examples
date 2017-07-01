package mathapp

import "math"

// defining a function with capital letter makes it automatically exported.
// yeah, should just use the one from math package instead, just learning here :)
func Sqrt(x float64) float64 {
	//just learning so using go imple :)

	return math.Sqrt(x)
}