package sprint

func ReverseAlphabetValue(ch rune) rune {
	// sets an if condition that checks if the rune value of ch is lower than the rune value of a or higher than the rune value of z
	if ch < 'a' || ch > 'z' {
		// when the condition is met, returns the value of ch
		return ch
	}
	// if the if condition is not met returns 219 (the sum of the integer values of a and z) minus the value of ch
	return 219 - ch
}