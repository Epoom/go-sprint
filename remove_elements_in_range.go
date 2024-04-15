package sprint /*main*/

/*import "fmt"*/

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
    length := len(arr)

    // Normalize indices to fall within the slice bounds
    for from < 0 {
        from += length
    }
    for to < 0 {
        to += length
    }

    for from >= length {
        from -= length
    }
    for to >= length {
        to -= length
    }

    if from > to {
        from, to = to, from
    }
    result := append(arr[:from], arr[to:]...)
    
    if len(result) == 0 {
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