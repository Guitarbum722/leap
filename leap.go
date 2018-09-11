package leap

import (
	"strconv"
)

// IsLeapYear returns the provided year and 1 if the year is a leap year; otherwise 0
func IsLeapYear(year int) string {
	isLeap := 48
	if year%400 == 0 {
		isLeap = 49
	} else if year%100 == 0 {
		isLeap = 48
	} else if year%4 == 0 {
		isLeap = 49
	}

	return strconv.Itoa(year) + " - " + string(isLeap)
}
