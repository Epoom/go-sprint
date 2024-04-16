package sprint

func StrConcatWith(strs []string, sep string) string {

newstr := ""
for i, str := range strs {
	newstr += str
	if i < len(strs)-1 {
		newstr += sep
	}
}
return newstr
}


/*Create a Go function that takes a slice of strings and a separator string as its parameters. 
The function should return a new string by concatenating all the strings in the slice, with each 
string separated by the provided separator.*/