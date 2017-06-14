package leap

const testVersion = 3

func IsLeapYear(y int) bool {
	// Write some code here to pass the test suite.
	var isLeapYear = false

	if y%4 == 0 {
		if y%100 == 0 {
			isLeapYear = false
			if y%400 == 0 {
				isLeapYear = true
			}
		} else {
			isLeapYear = true
		}
	}

	return isLeapYear
}
