package sprint

func GetLastRune(s string) rune {
var last rune
for _, c := range s {
		last = c
	}
return last
}