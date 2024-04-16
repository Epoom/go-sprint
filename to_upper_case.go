package sprint

func ToUpperCase(s string) string {
	var str string
	str = ""
	for _, c := range s {
		if c > 'a' || c < 'z' {
			c -= 32
			str = str + string(rune(c))
		}
		str = str + string(rune(c))
	}
	return str
}