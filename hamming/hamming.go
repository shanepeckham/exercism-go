package hamming

import (
	"errors"
)

const testVersion = 6

func Distance(a, b string) (int, error) {

	var hammingInt int

	if err := len(a) != len(b); err == true {
		return -1, errors.New("String values are of unequal length")
	}

	sliceA := a[:len(a)]
	sliceB := b[:len(b)]

	for i := 0; i < len(a); i++ {
		if sliceA[i:i+1] != sliceB[i:i+1] {
			hammingInt++
		}

	}
	return hammingInt, nil

}
