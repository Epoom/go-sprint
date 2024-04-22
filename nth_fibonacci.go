package sprint

func NthFibonacci(index int) int {

res := 0

if index < 0 {
	return -1
} else if index == 0 {
	return 0
} else if index == 1 {
	return 1
} else {
	res = NthFibonacci(index -1) + NthFibonacci(index -2)
}
return res
}