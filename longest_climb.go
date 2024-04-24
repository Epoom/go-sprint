package sprint 

func LongestClimb(arr []int) []int {
start := 0
maxLen := 1
currentLen := 1

for i := 1; i < len(arr); i++ {
	if arr[i] > arr[i-1] {
		currentLen++

		if currentLen > maxLen {
			maxLen = currentLen
			start = i - maxLen + 1
		}
	} else {
		currentLen = 1
	}
}
return arr[start : start+maxLen]
}