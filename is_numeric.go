package sprint

func IsNumeric(s string) bool {
	var valid bool
	valid = true
	for _, c := range s {
		if c < '0' || c > '9' {
			valid = false
			return valid
		}
	 
	}
	return valid
}
