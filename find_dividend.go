package sprint

func FindDividend(from, to, divisor int) int{
	// sets a for loop condition that first changes the value of i to the value of from
	// then checks if the value of i is lower than the value of to, if so, it elevates the value of i by 1
for i := from ; i < to ; i++ {
	// sets a if condition that checks if i modulo divisor equals 0
	if i % divisor == 0 {
		// when the if condition is met, returns the value of i
		return i
	}
}	
// if the if condition is not met by the time the for loop ends, returns -1
	return -1
}