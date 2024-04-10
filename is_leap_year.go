package sprint

func IsLeapYear(year int) bool {
	if year % 4 == 0 {
		if year % 400 != 0 {
			return bool(false)
			}
		return bool(true)
	}
	return bool(false)
}