package sprint

import "fmt"

func Combinations() string {
	str := ""
	for a := 0; a <= 7; a++ {
		for b := a+1; b <= 8; b++ {
			for c := b+1; c <= 9; c++ {
				str += fmt.Sprintf("%d%d%d, ", a, b, c)
			}
		}
	}
	newstr := ""
	newstr = str[:len(str)-2]
	return newstr
}

