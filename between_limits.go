package sprint

func BetweenLimits(from, to rune) string {
	// starts an if condition that checks if the value of to is lower than the value of from
	if to < from {
		//when the if condition is met it switches the values of to and from
		to, from = from, to
	}
	// create a variable str and set is value to "" aka an empty string
	str := ""
	//create a for loop condition that first sets the value of i to be the value of from + 1
	//it then checks if the value of i is lower than the value of to, if it is, it elevates the value of i by 1
		for i := from + 1; i < to; i++ {
			// if the value of i has been elevated by 1, it sets the value of str to be its current value + the string value of i
			str = str + string(i)
		}
		// once the for loop end, it return the value of str as a string
		return string(str)
}