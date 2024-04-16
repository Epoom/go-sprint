package sprint

/*import "fmt"*/

func ToCapitalCase(s string) string {
capitalize := true
var str string

for _, i := range s {
	 if i >= 'A' && i <= 'Z' {
		if !capitalize {
			str += string(i + 32) 
		} else {
			str += string(i)
		}
	} else if i >= 'a' && i <= 'z' {
		if capitalize {
			str += string(i - 32)
		} else {
			str += string(i)
		}
		capitalize = false
	} else if i < 'A' || i > 'z' || (i < 0 && i > 9) {
		capitalize = true
		str += string(i)
	} else {
		str += string(i)
	}
}
return str
}



/*func main() {
	str := "Hello! How are you? How+are+things+4you?"
	fmt.Println(ToCapitalCase(str))
}*/