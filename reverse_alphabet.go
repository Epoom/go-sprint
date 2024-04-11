package sprint

func ReverseAlphabet(step int) string {
	str := ""
	if step <= 0 {
		step = 1
	}
	for i := 122; i >= 97; i -= step {
		str = str + string(rune(i))
	}
	return string(str)
}