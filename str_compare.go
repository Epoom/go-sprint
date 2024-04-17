package sprint

func StrCompare(a, b string) int {

if a == b {
	return 0
}
if a < b {
	return -1
}
return 1

}

/*Create a Go function that mimics the behavior of the Compare function. It takes two strings, a and b, as input and returns an integer.
The function should compare the two strings and return:

0 if the strings are equal,
-1 if a comes before b in lexicographic order,
1 if a comes after b in lexicographic order.*/