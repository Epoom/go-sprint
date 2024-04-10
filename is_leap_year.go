package sprint

func IsLeapYear(year int) bool {
	if year % 4 == 0 {
		return bool(true)
	}
	if year % 100 == 0 {
		if year % 400 != 0{
			return bool(false)
		}
	}
	return bool(false)
}