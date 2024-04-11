package sprint

import "fmt"

func Pairs() string {
	str := ""
for i := 0; i <= 99; i++ {
	for j := i + 1; j <= 99; j++ {
		str += fmt.Sprintf("%02d %02d, ", i, j)
			if i == 98 && j == 99 {
				str += fmt.Sprintf("%02d %02d", i, j)
		}
	}

	}
	return str
}

