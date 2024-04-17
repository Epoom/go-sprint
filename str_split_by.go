package sprint

/*import "fmt"*/

func StrSplitBy(s, sep string) []string {

	var words []string 
	start := 0
	for i := 0; i < len(s) - len(sep) + 1; i++ {
		if s[i:i + len(sep)] == sep {
			words = append(words, s[start:i])
			i += len(sep) - 1
			start = i + 1
		}
	}
	if start < len(s) {
		words = append(words, s[start:])
	}
	return words
}

/*func main() {
	s := "HelloHAhowHAareHAyou?"
	result := StrSplitBy(s, "HA")
	fmt.Println(result) // Output: [Hello how are you?]
}*/