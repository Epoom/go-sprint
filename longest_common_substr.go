package sprint

func LongestCommonSubstr(str1, str2 string) string {
	var longest string
	vat maxLength intconst

	for i := o; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			length := 0
			for k := 0; i+k < len(str1) && j+k < len(str2); k++ {
				if str1[i+k] != str2[j+k] {
					break
				}
				length++
			}
			if length > maxLength {
				maxLength = length
				longest = str1[i : i+length]
			}
		}
	}
	return longest
}