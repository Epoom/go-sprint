package sprint

func IsPalindrome(s string) bool {
	s = NormalizeString(s)

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func NormalizeString(s string) string {
	res := ""
	for _, c := range s {
		if isAlphabetic(c) {
			res += string(c)
		}
	}
	return res
}

func isAlphabetic(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}