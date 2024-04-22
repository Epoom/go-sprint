package sprint

func IsPrime(n int) bool {

	if n < 2 {
		return false
	}
	for i := 2; i < 46340 && i * i <= n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}