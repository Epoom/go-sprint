package sprint

func ShiftBy(r rune, step int) rune {
	// sets an if condition that checks if the value of r is higher than a or lower than z
	if r < 'a' || r > 'z' {
		// when the if condition is met, return the value of r
		return r
	}
	// return the value of r minus the ascii value of a, plus the ascii value of step(user inputted number)
	// then it takes that new value and modulos it by 26 (the number of letters in the slphabet) + the ascii value of a
	// therefore shifting the alphabet by the value of step
	return ((r - 'a' + rune(step)) % 26) + 'a'
} 
