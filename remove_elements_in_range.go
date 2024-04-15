package sprint /*main*/

/*import "fmt"*/

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
	
	for from < 0 {
		from += len(arr)
	}
	from = from % len(arr)

	for to < 0 {
		to += len(arr)
	}
	to = to % len(arr)

	if from > to {
		from, to = to, from
	}
	result := append(arr[:from], arr[to:]...)
	if result == nil {
		return []float64{}
	}

	return result
}

/*func main() {
	input := []float64{10., .8, -.4, 20., 7.7, 3.}
	from := 1
	to := -36

	result := removeElementsInRange(input, from, to)
	fmt.Println(result) // Output: [10 7.7 3]
}*/