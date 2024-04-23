package sprint

func IsSorted(f func(a, b string) int, arr []string) bool {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		if f(arr[i], arr[i+1]) > 0 {
			return false
		}
	}
	return true
}

func StrCompare(a, b string) int {

	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
	
	}