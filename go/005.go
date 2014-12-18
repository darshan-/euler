package main

import "fmt"

type intMap map[int64] int64

type Colector struct {
	existing intMap
}

func (col *Colector) ensure(ensured intMap) {
	for key, value := range ensured {
		cur := col.existing[key]
		if cur < value {
			col.existing[key] = value
		}
	}
}

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

var primer Primer

func factors(n int64) intMap {
	primer.reset()
	f := intMap {}
	for {
		i := primer.nextPrime()
		for n % i == 0 {
			f[i] += 1
			if n == i { return f }
			n /= i
		}
	}

	panic("Unreachable"); // Please the compiler
}

const GOAL = 20

func main() {
	col := new(Colector)
	col.existing = intMap {}

	for i := int64(2); i <= GOAL; i++ {
		col.ensure(factors(i))
	}

	tot := int64(1)
	for key, value := range col.existing {
		for i := int64(0); i < value; i++ {
			tot *= key
		}
	}

	fmt.Println(tot)
}
