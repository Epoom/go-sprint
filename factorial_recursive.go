package sprint

func FactorialRecursive(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * FactorialRecursive(n -1)
	}
}