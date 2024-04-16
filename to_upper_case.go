package sprint

func ToUpperCase(s string) string {
	var str string
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			str += string(c - 32)
		} else {
			str += string(c)
		}
	}
	return str
}