package diffsquares

const testVersion = 1

func SquareOfSums(nati int) int {

	var squareofsums int
	for i := nati; i > 0; i-- {
		squareofsums += i
	}
	return squareofsums * squareofsums
}

func SumOfSquares(nati int) int {

	var sumofsquares int
	for i := nati; i > 0; i-- {
		sumofsquares += i * i
	}
	return sumofsquares

}

func Difference(nati int) int {

	sos := SquareOfSums(nati)
	sus := SumOfSquares(nati)

	return sos - sus
}
