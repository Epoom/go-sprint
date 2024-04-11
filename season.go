package sprint


func Season(month string) string {
	// set a switch condition to the string value of month
	switch month {
		// if the string value of month is "jan", "feb" or "dec"
	case "jan", "feb", "dec":
		// return the string "winter"
		return "winter"
		// if the string value of month is "mar", "apr" or "may"
	case "mar", "apr", "may":
		// return the string "spring"
		return "spring"
		// if the string value of month is "jun", "jul" or "aug"
	case "jun", "jul", "aug":
		// return the string "summer"
		return "summer"
		// if the string value of month is "sep", "oct" or "nov"
	case "sep", "oct", "nov":
		// return the string "autumn"
		return "autumn"
		// if none of the cases of the switch are met, then goes to the default return
	default:
		// default return is set to the strings "invalid input: " plus the string value of month
		return ("invalid input: " + month)

	}
}
