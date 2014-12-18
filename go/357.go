package main

const GOAL = 100 * 1000 * 1000

var numbers [GOAL + 2]bool

func satisfiesReq(n int) bool {
	for i := 1; i * i <= n; i++ {
		if (n % i == 0) && numbers[i + n / i] {
			return false
		}
	}
	return true
}

func main() {
	for i := 2; i * i < GOAL + 2; {
		for j := i * i; j < GOAL + 2; j += i {
			numbers[j] = true
		}
		for i++; numbers[i]; i++ {}
	}

	var sum int64

	for i := 1; i <= GOAL; i++ {
		if satisfiesReq(i) {
			sum += int64(i)
		}
	}

	println(sum)
}
