package sprint

func AdvancedSortWordArr(a []string, f func(a, b string) int) []string {
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if !(a[j], a[j+1]) {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
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