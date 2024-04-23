package sprint

func ArrCountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, s := range tab {
		if f(s) {
			count++
		}
	}
	return count
}

func IsUpper(s string) bool {
	var valid bool
	valid = true
	for _, c := range s {
		if c < 'A' || c > 'Z' {
			valid = false
			return valid
		}
	 
	}
	return valid
}

func IsAlphanumeric(s string) bool {
	for _, c := range s {
        if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') && (c < '0' || c > '9') {
            return false
        }
    }
    return true
}

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