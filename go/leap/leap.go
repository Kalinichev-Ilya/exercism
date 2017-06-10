package leap

const testVersion = 3

func IsLeapYear(year int) bool {
	var result bool
	if year%4 != 0 {
		result = false
	} else if year%100 != 0 || year%400 == 0 {
		result = true
	}
	return result
}
