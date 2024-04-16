package sprint

func ToUpperCase(s string) string {
	var str string
	for i := 0; i < len(s); i++ {
		if s[i] > 97 || s[i] < 122 {
			str = str + string(rune(s[i] - 32))
		}
		str = str + string(rune(s[i]))
	}
	return str
}