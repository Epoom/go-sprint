package sprint

func IntVsFloat(i int, f float32) string {
	// set an if condition that checks if the value of f is equal to the value of i when i is converted from an integer to a float
	if f == float32(i) {
		// if the condition is met, returns the string "same", indicating that their values are the same
		return string("Same")
	}// set an if condition that checks if the value of f is lower than the value of i when i is converted from an integer to a float
	if f < float32(i) {
		// if the condition is met, returns the string "Integer", indicating that the value of the integer is higher
		return string("Integer")
	}
	// if neither if condition is met, returns string "Float", indicating that the value of the Float32 is higher
		return string("Float")
}