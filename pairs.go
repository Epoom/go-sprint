package sprint

import "fmt"

func Pairs() string {
	// declare a variable called str and make its value an empty string
	str := ""
	// start a for loop that first sets the variable i to the value of 0
	// then it checks if the value of i is below 99, if it is it increments it
	for i := 0; i <= 99; i++ {
		// when the first for loop increments i, it will move on to the second for loop
		// the second for loop first sets the variable j to the value of 1 so that j and i never have the same number since j is always 1 ahead
		// it then checks if j is below 99, if it is it increments j by 1
		for j := i + 1; j <= 99; j++ {
			// when j is incremented it goes on to this declaration
			// the string str will be written into by the sprintf function
			// %02d explanation: 2d means 2 digits, so it writes the incoming value as 2 digits, the 0 means that if the value given is not 2 digits long, it will occupy the first space with a 0
			str += fmt.Sprintf("%02d %02d, ", i, j)
	}
	}
	// declare a variable called newstr and make its value an empty string
	newstr := ""
	// declare the value of newstring to be the string value of str - its lenght by 2, therefore cutting out the last 2 runes from the string
	newstr = str[:len(str)-2]
	// returns the new string
	return newstr
}