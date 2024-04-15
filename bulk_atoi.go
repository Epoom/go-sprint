package sprint /*main*/

/*import "fmt"*/

func BulkAtoi(arr []string) []int {
var result []int
for i := 0; i < len(arr); i++ {
	num := StrToInt(arr[i])
	result = append(result, num)
	}
	return result
}


func StrToInt(s string) int {
	x := 1
	nb := 0
	i := 0
	if len(s) == 0 {
		return 0
	}
	for s[i] == '-' || s[i] == '+' {
		if s[i] == '-' {
			x *= -1
		}
		i++
	if s[i] == '-' || s[i] == '+' {
		return 0
	}

	}
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		nb = (nb * 10) + (int(s[i] - '0'))
		i++
	}
	if i < len(s) {
		return 0
	}
	nb *= x
	return nb
}

/*func main() {
	arr := []string{"8", "kood", "-13"}
	fmt.Println(BulkAtoi(arr))
}*/