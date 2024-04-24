package sprint

import (
	"sort"
	"strings"
)

func AreAnagrams(str1, str2 string) bool {
	 str1 = NormalizeString(str1)
	 str2 = NormalizeString(str2)

	 if len(str1) != len(str2) {
		return false
	 }
	 Runes1 := []rune(str1)
	 Runes2 := []rune(str2)

	 sort.Slice(Runes1, func(i, j int) bool {return Runes1[i] < Runes1[j]})
	 sort.Slice(Runes2, func(i, j int) bool {return Runes2[i] < Runes2[j]})
	
	 return string(Runes1) == string(Runes2)
}

func NormalizeString(s string) string {
	normalized := ""
	for _, char := range strings.ToLower(s) {
		if char >= 'a' && char <= 'z' {
			normalized += string(char)
		}
	}
	return normalized
}