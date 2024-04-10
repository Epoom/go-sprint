package sprint

func BetweenLimits(from, to rune) string {
	if to < from {
		to, from = from, to
	}
	str := ""
		for i := from + 1; i < to; i++ {
			
			str = str + string(i)
		}
		return string(str)
}