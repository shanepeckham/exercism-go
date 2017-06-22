package triangle

import (
	"math"
)

const testVersion = 3

type Kind string

const (
	NaT Kind = "NaT"
	Equ Kind = "Equ"
	Iso Kind = "Iso"
	Sca Kind = "Sca"
)

func CheckIsTriangle(a, b, c float64) bool {

	var result = true

	if math.Mod(a, a) != 0 || math.Mod(b, b) != 0 || math.Mod(c, c) != 0 {
		result = false
	}

	if a+b+c <= 0 {
		result = false
	}

	if a+b < c {
		result = false
	}

	if a+c < b {
		result = false
	}

	if b+c < a {
		result = false
	}

	return result
}

func KindFromSides(a, b, c float64) Kind {

	var result Kind

	if a != b && b != c {
		result = Sca
	}

	if (a == b || b == c) || (a == b || a == c) {
		result = Iso
	}

	if a == b && a == c {
		result = Equ
	}

	if !CheckIsTriangle(a, b, c) {
		result = NaT
	}

	return result
}
