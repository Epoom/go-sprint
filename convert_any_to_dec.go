package sprint

/*import (
	"fmt"
	"math"
	"strings"
)*/

func ConvertAnyToDec(s string, base string) int {
	

	if len(base) < 2 || strings.ContainsAny(base, "+-") {
		return 0
	}
	
	baseMap := [256]int{}
	for i, c := range base {
		baseMap[c] = i
	}

	result := 0
	for i := 0; i < len(s); i++ {
		charValue := baseMap[rune(s[i])]
		if charValue >= len(base) {
			return 0
		}
		result += charValue * int(math.Pow(float64(len(base)), float64(len(s)-1-i)))
	}
return result
}

/*func main() {
fmt.Println(ConvertAnyToDec("125", "0123456789")) // should output 125
fmt.Println(ConvertAnyToDec("1111101", "01")) // should output 125
fmt.Println(ConvertAnyToDec("7D", "0123456789ABCDEF")) // should output 125
fmt.Println(ConvertAnyToDec("uoi", "choumi")) // should output 125
fmt.Println(ConvertAnyToDec("bbbbbab", "-ab")) // should output 0
}*/

/*Create a Go function that takes two parameters:

s: a numeric string in a specific base.
base: a string containing unique characters representing the available digits in that base.
The function should return the integer value of s in the given base. If the base is not valid, it returns 0.
Here are the validity rules for the base:
The base must contain at least 2 unique characters.
The base should not contain the characters + or -.
The function only works with valid string numbers that consist of elements present in the base. It does not handle negative numbers.*/