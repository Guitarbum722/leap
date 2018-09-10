package leap

import "fmt"

// IsLeapYear returns true if the input year is a leap year
func leapYear(year int) string {
	isLeap := 0
	if year%400 == 0 {
		isLeap = 1
	} else if year%100 == 0 {
		isLeap = 0
	} else if year%4 == 0 {
		isLeap = 1
	}

	return fmt.Sprintf("%v - %v", year, isLeap)
}
