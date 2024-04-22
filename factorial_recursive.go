package sprint

func FactorialRecursive(n int) int {

	if n <= 0 {
		return 0
	} else {
		return n * FactorialRecursive(n -1)
	}
}