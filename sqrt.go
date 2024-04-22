package sprint

func Sqrt(n int) int {

	if n  < 0 {
		return 0
	}
	if n == 1 || n == 0 {
		return n
	}
	for i := 2; i <= n / i; i++ {
		if i * i == n {
			return i
		}
	}
	return 0
}