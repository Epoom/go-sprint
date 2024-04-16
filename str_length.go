package sprint

func StrLength(s string) []int {
RinS := 0
byteLength := len(s)

for range s {
	RinS++
}
return []int{RinS, byteLength}
}