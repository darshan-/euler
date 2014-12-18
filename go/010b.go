package main

const GOAL = 2 * 1000 * 1000
const SQRT = 1414

func main() {
	var numbers [GOAL]bool
	var sum int64

	for i := 2; i <= SQRT; {
		sum += int64(i)

		for j := i * i; j < GOAL; j += i {
			numbers[j] = true
		}

		for i++; numbers[i]; i++ {}
	}

	for i := SQRT + 1; i < GOAL; i++ {
		if ! numbers[i] {
			sum += int64(i)
		}
	}

	println(sum)
}
