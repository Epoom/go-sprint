package sprint

func IsSorted(f func(a, b string) int, arr []string) bool {
	n := len(arr)
	ascending := f(arr[0], arr[1]) > 0
    for i := 1; i < n-1; i++ {
        result := f(arr[i], arr[i+1])
        if (ascending && result < 0) || !ascending && result > 0  {
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