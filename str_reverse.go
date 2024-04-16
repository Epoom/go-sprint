package sprint

func StrReverse(s string) string {
runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; {
			runes[i], runes[j] = runes[j], runes[i]
			i++
			j--
		}
		return string(runes)
	}

