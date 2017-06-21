package pangram

import (
	"strings"
)

const testVersion = 1

func IsPangram(s string) bool {

	s = strings.ToLower(s)
	charMap := make(map[rune]int)

	for _, i := range s {

		if 'a' <= i && i <= 'z' {
			charMap[i] = 1
		}
	}

	return len(charMap) >= 26

}
