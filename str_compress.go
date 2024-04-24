package sprint

import "fmt"

func StrCompress(input string) string {
	 res := ""
	 count := 1

	 for i := 0; i < len(input); i++ {
		if i+1 < len(input) && input[i] == input[i+1] {
			count++
		} else {
			if count > 1 {
				res += fmt.Sprintf("%d%s", count, string(input[i]))
			} else {
				res += string(input[i])
			}
			count = 1
		}
	 }
	 return res
}