package sprint /*main*/

/*import "fmt"*/

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
	
	for from < 0 || from > len(arr) {
		if from < 0 {
			from += len(arr)
		}
		if from > len(arr) {
			from -= len(arr)
		}
	}
	for to < 0 || to > len(arr) {
		if to < 0 {
			to += len(arr)
		}
		if to >= len(arr) {
			to = len(arr)
		}
	}
	if from > to {
		from, to = to, from
	}
	if to == len(arr) {
		return arr[:from]
	}
	result := append(arr[:from], arr[to:]...)

	return result
}
/*func main() { y
	input := []float64{10., .8, -.4, 20., 7.7, 3.}
	from := 1
	to := -36

	result := removeElementsInRange(input, from, to)
	fmt.Println(result) // Output: [10 7.7 3]
}*/