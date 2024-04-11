package sprint



func Countdown(n int) string {
	str := ""
	for i := n; i > 0; i -= 2 {
		str += string('0' + byte(i)) + ", "
	}
	newstr := ""
	newstr = str + "0!"
	return newstr
}