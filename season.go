package sprint

import (
	"fmt"
)

func Season(month string) string {
	switch month {
	case "jan", "feb", "dec":
		fmt.Printf ("winter")
	case "mar", "apr", "may":
		fmt.Printf ("spring")
	case "jun", "jul", "aug":
		fmt.Printf ("summer")
	case "sep", "oct", "nov":
		fmt.Printf ("autumn")
	default:
		fmt.Printf("\"Invalid input: %s\"", month)
	}
	return ""
}
