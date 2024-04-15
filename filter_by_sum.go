package sprint

/*import "fmt"*/

func FilterBySum(arr [][]int, limit int) [][]int {
var newarr [][]int
for _, subarr := range arr {
	sum := 0
	for _, num := range subarr {
		sum += num
	}
if sum >= limit {
	newarr = append(newarr, subarr)
}
}
return newarr
}

/*func main() {
	arr := [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}
	limit := 9
	fmt.Println(FilterBySum(arr, limit))

}*/