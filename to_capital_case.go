package sprint

func ToCapitalCase(s string) string {
count := 0
var str string
for _, i := len(s) {
	count++
	if count > 1 {
		if i >= 'A' && i <= 'Z' {
			str += string(c + 32)
			}
		}
		str += string(c)
	}
	if count == 1 {
		if i >= 'a' && i <= 'z' {
			str += string(c - 32)
		}
		str += string(c)
	}
	if i == " " {
		count = 0
	}
	return str
}