package secret

import (
	"fmt"
	"math/big"
	"strconv"
)

const testVersion = 2

var verbs [4]string

func Handshake(n uint) []string {

	verbs[0] = "wink"
	verbs[1] = "double blink"
	verbs[2] = "close your eyes"
	verbs[3] = "jump"
	var secret [4]string
	var reverseSecret [4]string
	var reverse = true
	var a []string

	i64 := int64(n)

	ip := big.NewInt(i64)
	s := fmt.Sprintf("%b", ip)
	sh, err := strconv.Atoi(s)

	if err == nil {

		var j int
		var done = false

		for i := 0; done == false; {
			if sh >= 10000 {
				sh -= 10000
				reverse = false
			}
			if sh >= 1000 {
				secret[i] = verbs[3]
				sh -= 1000
				i++
			}
			if sh >= 100 {
				secret[i] = verbs[2]
				sh -= 100
				i++
			}
			if sh >= 10 {
				secret[i] = verbs[1]
				i++
				sh -= 10
			}
			if sh == 1 {
				secret[i] = verbs[0]
				i++
				sh--
			}
			if sh == 0 {
				j = i
				done = true
			}
			if sh < 0 {
				sh *= -1
			}
		}

		reverseSecret = secret

		itemCount := len(reverseSecret)
		for i := 0; i < itemCount/2; i++ {
			mirrorIdx := itemCount - i - 1
			reverseSecret[i], reverseSecret[mirrorIdx] = reverseSecret[mirrorIdx], reverseSecret[i]
		}

		a = secret[0:j]
		b := reverseSecret[len(reverseSecret)-j:]

		if reverse {
			return b
		}
	}

	return a
}
