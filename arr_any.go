package sprint

func ArrAny(f func(string) bool, a []string) bool {
	for _, s := range a {
		if f(s) {
			return true 
		}
	}
	return false
}

func IsUpper(s string) bool {
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
