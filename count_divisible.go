package sprint

func CountDivisible(from, to, step, divisor int) int {
	// set an if condition to check if the value of step is lower or equal to 0
	if step <= 0 {
		// if the condition is met, returns 0
		return 0
	}// set an if condition to check if the value of divisor is equal to 0
	if divisor == 0 {
// if the condition is met, return 0
		return 0
	}
	// create a variable called count and set its value to 0
	count := 0
	// create a for loop condition that sets the value of i to the value of from
	// it then checks if the value of i is lower than the value of to, elevates the value of i by the value of step
	for i := from ; i < to; i += step {
		// if the for loop condition is met, checks the if condition that if i module divisor equals 0
		if i % divisor == 0 {
			// when if conditions is met, sets the value of count to itself + 1, a simpler way to do it would be to write count++
			count = count + 1
		}
	} 
	// when for loop stops, returns the new value of count
	return count
}