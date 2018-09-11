package leap

// IsLeapYear returns the provided year and 1 if the year is a leap year; otherwise 0
func IsLeapYear(year int) string {
	isLeap := 0
	if year%400 == 0 {
		isLeap = 1
	} else if year%100 == 0 {
		isLeap = 0
	} else if year%4 == 0 {
		isLeap = 1
	}

	return string(year) + " - " + string(isLeap)
}
