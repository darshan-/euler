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
	for i = 100; ; i++ {
		if isBouncy(i) { count++ }
		if (int64(count) * 100 == int64(i) * 99) { break }
	}
	println(i)
}
