package sprint

func IsLeapYear(year int) bool {
	// set a if condition that checks if the value of year modulo 4 equals 0
	if year % 4 == 0 {
		// when the if condition is met, sets another if condition that checks if the value of year modulo 100 equals 0
		// AND the value of year modulo 400 does NOT equal 0 
		if year % 100 == 0 && year % 400 != 0 {
			// when this if condition is met, return the boolean value "false"
			return bool(false)
			}
		// when the first if condition is met and the second condition is NOT met, returns the boolean value "true"
		return bool(true)
	}
	// if the 1st if condition is not met, return the boolean value "false"
	return bool(false)
}