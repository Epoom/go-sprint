package sprint

func ArrMap(f func(int) bool, a []int) []bool {
 result := make([]bool, len(a))
 for i, val := range a {
	result[i] = f(val)
 }
 return result
}

func IsNegative(n int) bool {
	
	if n < 0 {	
	}
	return bool(false)
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