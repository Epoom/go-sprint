package sprint

func ToUpperCase(s string) string {
	var str string
	for i := 0; i < len(s); i++ {
		if i >= 97 && i <= 122 {
			str = str + string(rune(i - 32))
		} else {
		str = str + string(i)
		}
	}
	return str
}