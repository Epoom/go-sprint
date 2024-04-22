package sprint

func NextPrime(n int) int {

	if n < 2 {
		return 2
	}
	if IsPrime(n) {
		return n
	}
	next := n + 1
	for NextPrime(next) != 0 {
		n++
	}
	return next
}

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