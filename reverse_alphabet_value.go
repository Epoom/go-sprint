package sprint

func ReverseAlphabetValue(ch rune) rune {
	if ch < 'a' || ch > 'z' {
		return ch
	}

	return 'z' - (ch - 'a')
}