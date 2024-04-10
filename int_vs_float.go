package sprint

func IntVsFloat(i int, f float32) string {
	if f == float32(i) {
		return string("Same")
	}
	if f < float32(i) {
		return string("Integer")
	}
		return string("Float")
}