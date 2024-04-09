package sprint

import "strings"

func StrConcat(s1, s2, delim string) string {
	x := strings.Join([]string{s1, s2}, delim)	
	return x
}