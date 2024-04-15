package sprint

func AlphaNumber(n int) string {
letters := "abcdefghij"
str := ""
negative := false
if n < 0 {
negative = true
n = -n
}

if n == 0 {
	str = "a"
}

for n > 0 {
	// modulos the value of n by 10 aka grabbing the most right digit
	ch := n % 10
	// turning the value of str into the string value of the rune of the letter array + str
	str = string(letters[ch]) + str
	// changes the value of n to itself divided by 10 aka deleting the most right digit that was converted into a letter above
	n /= 10

}
if negative == true {
	str = "-" + str
}
return str
}

/*func main() {
	fmt.Println(AlphaNumber(0))
}*/