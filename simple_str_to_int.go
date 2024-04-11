package sprint

func SimpleStrToInt(s string) int {
	res := 0
	for _, digit := range s {
		res = res * 10 + int(digit - '0')
	}
	return res
}