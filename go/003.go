package main

import (
	"fmt"
	"time"
)

var t time.Time

var nextPrime = func() (func() int64) {
	cur    := int64(1)
	var primes []int64;

	isPrime := func(n int64) bool {
		for _, prime := range primes {
			if n % prime == 0 {
				return false
			}
		}
		return true
	}

	return func() int64 {
		for cur++; !isPrime(cur); cur++ {}
		primes = append(primes, cur)
		if (len(primes) % 1000 == 0) {
			fmt.Print(time.Since(t))
			fmt.Print("\tprimes now has ", len(primes))
			fmt.Println("\tlargest prime: ", primes[len(primes) - 1])
		}
		return cur
	}
}()

func maxPrimeFactor(n int64) int64 {
	for {
		i := nextPrime()
		for n % i == 0 {
			if n == i { return i }
			n /= i
		}
	}

	panic("Unreachable"); // Please the compiler
}

func main() {
	t = time.Now()
	//println(maxPrimeFactor(13195))
	println(maxPrimeFactor(600851475143))
}
