package sprint

func GenerateRange(min, max int) []int {

length := max - min
var a []int

for i := 0; i < length; i++ {
	a = append(a, min+i)
	}
	return a
}