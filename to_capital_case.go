package sprint

func ToCapitalCase(s string) string {
count := 0
var str string
for  i := 0; i <= len(s); i++ {
	count++
	if count > 1 {
		if i >= 'A' && i <= 'Z' {
			str += string(i + 32)
			}
		}
		str += string(i)
	if count == 1 {
		if i >= 'a' && i <= 'z' {
			str += string(i - 32)
		}
		str += string(i)
	}
	if i == ' ' {
		count = 0
	}
}
	return str
}