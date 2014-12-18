package main

type Primer struct {
	cur int
	primes []int
}

func (primer *Primer) isPrime(n int) bool {
	for _, prime := range primer.primes {
		if n % prime == 0 {
			return false
		}
		if prime * prime >= n {
			return true
		}
	}
	return true
}

func (primer *Primer) nextPrime() int {
	for primer.cur++; !primer.isPrime(primer.cur); primer.cur++ {}
	primer.primes = append(primer.primes, primer.cur)

	return primer.cur
}

func (primer *Primer) reset() {
	primer.cur = 1
}

func main() {
	var primer Primer
	primer.reset()

	var sum int64

	for p := 0; p < 2*1000*1000; p = primer.nextPrime() {
		sum += int64(p)
	}

	println(sum)
}

// 142913828922
