package sprint

func ToUpperCase(s string) string {
	var str string
	for i := 0; i < len(s); i++ {
		if i >= 'a' && i <= 'z' {
			str += string(i - 32)
		} else {
		str += string(i)
		}
	}
	return str
}