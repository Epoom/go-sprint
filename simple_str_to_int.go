package sprint

func SimpleStrToInt(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		// digit is assigned the character at index i in the string s
		digit := s[i]
		//int(digit - '0') converts the character digit to its numeric equivalent
		//by subtracting the ASCII value of '0' from the ASCII value of the character.
		res = res * 10 + int(digit - '0')
	}
	return res
}