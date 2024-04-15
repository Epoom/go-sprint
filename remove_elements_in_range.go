package sprint /*main*/

/*import "fmt"*/

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
	
	if from < 0 {
		from += len(arr)
	}
	from = from % len(arr)
	if to < 0 {
		to += len(arr)
	}

	to = to% len(arr)
	if from > to {
		from, to = to, from
	}
arr = append(arr[:from], arr[to:]...)
	if arr == nil {
		return []float64{}
	}

	return arr
}

/*func main() {
	input := []float64{10., .8, -.4, 20., 7.7, 3.}
	from := 1
	to := -36

	result := removeElementsInRange(input, from, to)
	fmt.Println(result) // Output: [10 7.7 3]
}*/