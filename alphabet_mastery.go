package sprint

func AlphabetMastery(n int) string {
	str := ""
	for i := 97; i < 97 + n; i++ {
		str = str + string(rune(i))
	}
	return string(str)
}