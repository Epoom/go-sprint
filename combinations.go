package sprint

import "fmt"

func Combinations() string {
	// declare a variable called str and make its value an empty string
	str := ""
	// start a for loop that first sets the variable a to the value of 0
	// then it checks if the value of a is below 7, if it is it increments it
	for a := 0; a <= 7; a++ {
		// when the first for loop increments a, it will move on to the second for loop
		// the second for loop first sets the variable b to the value of a+1 so that b and a never have the same number since b is always 1 ahead
		// it then checks if b is below 8, if it is it increments b by 1
		for b := a+1; b <= 8; b++ {
		// when the second for loop increments b, it will move on to the third for loop
		// the third for loop first sets the variable c to the value of b+1 so that c and b never have the same number since c is always 1 ahead
		// it then checks if c is below 9, if it is it increments b by 1
			for c := b+1; c <= 9; c++ {
			// when c is incremented it goes on to this declaration
			// the string str will be written into by the sprintf function
			// $d means that printf should write the integer value into the string
				str += fmt.Sprintf("%d%d%d, ", a, b, c)
			}
		}
	}
	// declare a variable called newstr and make its value an empty string
	newstr := ""
	// declare the value of newstring to be the string value of str - its lenght by 2, therefore cutting out the last 2 runes from the string
	newstr = str[:len(str)-2]
	// returns the new string
	return newstr
}