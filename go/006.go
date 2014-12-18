package main

func sumOfSquares(n uint64) (sum uint64) {
	for i := uint64(1); i <= n; i++ {
		sum += i * i
	}

	return sum
}

func squareOfSum(n uint64) (sum uint64) {
	for i := uint64(1); i <= n; i++ {
		sum += i
	}

	return sum * sum
}

func main() {
	n := uint64(100)

	println(squareOfSum(n) - sumOfSquares(n))
}
