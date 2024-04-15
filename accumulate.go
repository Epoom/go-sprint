package sprint

func Accumulate(n int) int {
	//create a variable called sum and set its value to 0
	sum := 0
	// start a if condition where it checks IF the integer n is higher or equal to 0
	if  n >= 0 {
		// when the if condition is met it starts the for loop
		// the for loop condition is set so that the value of "i" is set to 0
		// it then compares the value of "i", and if it is lower or equal to the value of "n", it elevates the value of "i" by 1
		for  i := 0 ; i <= n; i++  {
			//if the value of i has increased, it adds the value of i to the value of sum, therefore giving the sum of those values
			sum += i
		}
		// once the value of i equals the value of n, the for loops stops and in returns the new value of sum
		return sum
	}
	return 0
	//when the if condition is not met it skips the for loop and immediately returns 0
}