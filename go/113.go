package main

const (
	DESC = iota
	UNKNOWN
	ASC
)

func isBouncy(n int) bool {
	last_digit := -1
	direction  := UNKNOWN

	for n > 0 {
		digit := n % 10

		if last_digit >= 0 {
			if digit > last_digit {
				if direction == DESC {
					return true
				} else {
					direction = ASC
				}
			} else if digit < last_digit {
				if direction == ASC {
					return true
				} else {
					direction = DESC
				}
			}
		}

		n -= digit
		n /= 10
		last_digit = digit
	}

	return false
}

func main() {
	count := 0
	var i int
	for i = 100; i < 100 * 1000 * 1000; i++ {
		if isBouncy(i) { count++ }

		/*if i + 1 ==             1000 { println(i - count) }
		if i + 1 ==        10 * 1000 { println(i - count) }
		if i + 1 ==       100 * 1000 { println(i - count) }
		if i + 1 ==      1000 * 1000 { println(i - count) }
		if i + 1 == 10 * 1000 * 1000 { println(i - count) }*/
	}

	println(i - 1 - count)
}
