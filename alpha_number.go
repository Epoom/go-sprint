package sprint



func AlphaNumber(n int) string {
letters := "abcdefghij"
str := ""
negative := false
if n < 0 {
negative = true
n = -n
}

for n > 0 {
	ch := n % 10
	str = string(letters[ch]) + str
	n /= 10

}
if negative == true {
	str = "-" + str
}
return str
}

