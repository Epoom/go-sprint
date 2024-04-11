package sprint
// import the package "math" to use later on for mathematical purposes
import "math"

func Casting(n float64) int {
	// sets the value of x the rounded up version of the value of n with the Round function from the "math" package
	x := math.Round(n)
	// then returns the value of x as an integer
	return int(x)
}