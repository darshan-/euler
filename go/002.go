package main

var nextFib = func() (func() int) {
	a, b := 0, 1

	return func() int {
		a, b = b, a + b
		return b
	}
}()

const MAX = 4 * 1000 * 1000

func main() {
	sum := 0

	for {
		f := nextFib()

		if f > MAX {
			break
		}

		if f % 2 == 0 {
			sum += f
		}
	}

	println(sum)
}