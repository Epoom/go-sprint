package sprint

func StrToInt(s string) int {
	x := 1
	nb := 0
	i := 0
 	// basically checks for an empty string, and if its an empty string, returns 0
	if len(s) == 0 {
		return 0
	}
	// checks if the current character in the string is a - or a +
	for s[i] == '-' || s[i] == '+' {
		// then checks again, if it is a - it multiplies the value of x by -1, turning it into a - integer
		if s[i] == '-' {
			x *= -1
		}
		i++
	// checks if the next character in the string is also a + or -, if so, returns 0
	if s[i] == '-' || s[i] == '+' {
		return 0
	}

	}
	// for loop condition checks if i is lower thann the lenght of the string
	// and checks if the current value of i is between 0 and 9 aka a number
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		// multiplying nb by 10, effectively moving all digits in the string to the left
		// int(s[i] - '0' converts the elemnt value of i in the string to its integer value
		nb = (nb * 10) + (int(s[i] - '0'))
		i++
	}
	// if i i still lower than the lenght of the string after the for loop, it means there are non numbers in the string and it returns 0
	if i < len(s) {
		return 0
	}
	// multiplies our new integer by x to turn it into either a positive or negative number and returns it
	nb *= x
	return nb
}