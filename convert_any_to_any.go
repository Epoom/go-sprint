package sprint

import (
	"strings"
	"math"
)

func ConvertAnyToAny(nbr, baseFrom, baseTo string) string {

	if len(baseFrom) < 2 || strings.ContainsAny(baseFrom, "+-") {
		return ""
	}
	baseToMap := make(map[rune]int)
	for i, c := range baseFrom {
		baseToMap[c] = i
	}
	result := 0
	for i := 0; i < len(nbr); i++ {
		charValue := baseToMap[rune(nbr[i])]
		if charValue >= len(baseFrom) {
			return ""
		}
		result += charValue * int(math.Pow(float64(len(baseFrom)), float64(len(nbr)-1-i)))
	}
	var resStr string
	for result > 0 {
		digit := result % 10
		resStr = string(digit + '0') + resStr
		result /= 10
	}
return resStr
}

/*func main() {
fmt.Println(ConvertAnyToAny("101011", "01", "0123456789"))
}*/
/*Create a Go function that takes three string arguments: nbr representing a numeric value in a specific base baseFrom, 
and baseTo representing the desired base for the returned value. 
The function should convert the nbr from baseFrom to baseTo and return the result as a string.*/