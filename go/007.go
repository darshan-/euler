package main

type Primer struct {
	cur int64
	primes []int64
}

func (primer *Primer) nextPrime() int64 {
	isPrime := func(n int64) bool {
		for _, prime := range primer.primes {
			if n == prime {
				return true
			}
			if n % prime == 0 {
				return false
			}
		}
		return true
	}

	for primer.cur++; !isPrime(primer.cur); primer.cur++ {}
	primer.primes = append(primer.primes, primer.cur)

	return primer.cur
}

func (primer *Primer) reset() {
	primer.cur = 1
}

func main() {
	var primer Primer
	primer.reset()
	for i := 1; i < 10001; i++ { primer.nextPrime() }
	println(primer.nextPrime())
}