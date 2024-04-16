package sprint

func IsLower(s string) bool {
	var valid bool
	valid = true
	for _, c := range s {
		if c < 'a' || c > 'z' {
			valid = false
			return valid
		}
	 
	}
	return valid
}
