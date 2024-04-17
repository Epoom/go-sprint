package sprint

/*import "fmt"*/

func NbrBase(n int, base string) string {

var result string
var modulo int
thing := make(map[rune]bool)

for _, c := range base {
	if rune(c) == '+' || rune(c) == '-' {
		return "NV"
	}
	
	if thing[c] {
		return "NV"
	}
	thing[c] = true
}

if len(base) < 2 {
	return "NV"
}
isNegative := false
if n < 0 {
	n = n*-1
	isNegative = true
}

for n  > 0 {
	modulo = n % len(base)
	n = n / len(base)
	result = string(base[modulo]) + result
}
if isNegative {
	result = "-" + result
}

return result
}

/*func main() {
	fmt.Println(NbrBase(125, "01"))
}*/

/*Create a Go function that takes an integer n and a string base as parameters. The function should return the integer n represented in the specified base as a string.
Here are the validity rules for the base:

The base must contain at least 2 unique characters.
The base should not contain the characters + or -.
The function should handle negative numbers as well (see examples in the usage section). If the base is not valid, the function should return "NV" (Not Valid).*/