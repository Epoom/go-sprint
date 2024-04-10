package sprint

func Accumulate(n int) int {
	sum := 0
	if  n >= 0 {
		for  i := 0 ; i <= n; i++  {
			sum += i
		}
		return sum
	}
	return 0
}