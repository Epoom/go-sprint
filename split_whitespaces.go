package sprint

func SplitWhitespaces(s string) []string {

var words []string
current := ""
for _, c := range s {
	if c == ' ' || c == '\t' || c == '\n' {
		if current != "" {
			words = append(words, current)
			current = ""
		}
	} else {
		current += string(c)
	}
}
if current != "" {
	words = append(words, current)
}
return words
}