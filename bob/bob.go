package bob

import "strings"

const testVersion = 3

func Hey(s string) string {

	var response = "Whatever."
	s = strings.TrimSpace(s)
	upper := strings.ToUpper(s)

	if strings.ContainsAny(s, "?") && strings.LastIndex(s, "?") == len(s)-1 {
		response = "Sure."
	}

	if s == upper && strings.ContainsAny(upper, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		response = "Whoa, chill out!"
	}

	if len(s) == 0 || len(strings.Replace(s, "\t", "", -1)) == 0 {
		response = "Fine. Be that way!"
	}

	return response
}
