package main

import "math"

func palindromeOf(n int) int {
	var digits []int
	dup := n

	for dup > 0 {
		digits = append(digits, dup % 10)
		dup /= 10
	}

	for _, digit := range digits {
		n *= 10
		n += digit
	}

	return n
}

func largestNDigitNum(n int) int {
	return int(math.Pow10(n)) - 1
}

func smallestNDigitNum(n int) int {
	return int(math.Pow10(n - 1))
}

func isProductOfNDigits(num, nDigits int) bool {
	max := largestNDigitNum(nDigits)
	min := smallestNDigitNum(nDigits)

	for i := max; i > min; i-- {
		if num % i == 0 && num / i < max {
			return true
		}
		// shortcuts to return false?
	}

	return false
}

func largestPalOfNDigits(nDigits int) int {
	for i := largestNDigitNum(nDigits); i > 0; i-- {
		pal := palindromeOf(i)
		if isProductOfNDigits(pal, nDigits) {
			return pal
		}
	}

	panic("Compiler satisfaction")
}

func main() {
	println(largestPalOfNDigits(3))
}
