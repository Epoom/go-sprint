package sprint

func RemoveDuplicates(arr []int) []int {
	uniqueElem := make(map[int]bool)
	res := []int{}

	for _, num := range arr {
		if !uniqueElem[num] {
			uniqueElem[num] = true
			res =  append(res, num)
		}
	}
	return res
}