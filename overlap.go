package sprint

func Overlap(arr1, arr2 []int) []int {
	common := make([]int, 0)
	frequency := make(map[int]int)

	for _, num := range arr1 {
		frequency[num]++
	}

	for _, num := range arr2 {
		if frequency[num] > 0 {
			common = append(common, num)
			frequency[num]--
		}
	}
	common = Sort(common)
	return common
}

func Sort(arr []int) []int {
    for i := 0; i < len(arr); i++ {
        for j := i + 1; j < len(arr); j++ {
            if arr[i] > arr[j] {
                arr[i], arr[j] = arr[j], arr[i]
            }
        }
    }
    return arr
}