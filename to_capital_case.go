package sprint

/*import "fmt"*/

func ToCapitalCase(s string) string {
capitalize := true
var str string

for _, i := range s {
	if (i >= 'A' && i <= 'Z') || (i >= 'a' && i <= 'z') {
		if capitalize {
			if i >= 'a' && i <= 'z' {
				str += string(i - 32)
			} else {
				str += string(i)
			}
			capitalize = false
		} else {
			if i >= 'A' && i <= 'Z' {
				str += string(i + 32) 
			} else {
				str += string(i)
			}
		} 
	} else if i >= '0' && i <= '9' {
		str += string(i)
		capitalize = false
	} else {
		str += string(i)
		capitalize = true
	}
}
return str
}



/*func main() {
	str := "Hello! How are you? How+are+things+4you?"
	fmt.Println(ToCapitalCase(str))
}*/