package sprint

func GenerateRange(min, max int) []int {

lenght := max - min + 1
var a []int

for i := 0; i < lenght; i++ {
	a = append(a, min+i)
	}
	return a
}