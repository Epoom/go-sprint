package sprint

/*import "fmt"*/

func BalanceOut(arr []bool) []bool {
trueCount, falseCount := 0, 0

for _, val := range arr {
	if val {
		trueCount++
	} else {
		falseCount++
	}
}

diff := trueCount - falseCount

if diff == 0 {
	return arr
}
var valueToAdd bool
if diff <= 0 {
	valueToAdd = true
	diff = -diff
} else {
	valueToAdd = false
}

	for i := 0; i < diff; i++ {
	arr = append(arr, valueToAdd)
}
return arr
}

/*func main() {
	// Test the function with an example array
	arr := []bool{true, true, false, false, false}
	fmt.Println(BalanceOut(arr))
}*/

