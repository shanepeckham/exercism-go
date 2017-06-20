package acronym

import "strings"

const testVersion = 3

func Abbreviate(s string) string {

	sliceA := s[:len(s)]
	var t string
	for i := 0; i < len(s); i++ {

		if i == 0 {
			t += strings.ToUpper(sliceA[i : i+1])
		}
		if sliceA[i:i+1] == " " || sliceA[i:i+1] == "-" {
			t += strings.Replace(strings.ToUpper(sliceA[i:i+2]), " ", "", -1)
		}

	}
	return strings.Replace(t, "-", "", -1)
}
