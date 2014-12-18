package main

import (
	"fmt"
	"math/big"
	"time"
)

var bigInt big.Int // To call methods on
var t time.Time

var nextPrime = func() (func() *big.Int) {
	cur := big.NewInt(1)
	var primes []*big.Int;

	isPrime := func(n *big.Int) bool {
		for _, prime := range primes {
			if bigInt.Rem(n, prime).Cmp(big.NewInt(0)) == 0 {
				return false
			}
		}
		return true
	}

	return func() *big.Int {
		for cur.Add(cur, big.NewInt(1)); !isPrime(cur); cur.Add(cur, big.NewInt(1)) {}
		primes = append(primes, big.NewInt(0).Set(cur))
		if (len(primes) % 1000 == 0) {
			fmt.Print(time.Since(t))
			fmt.Print("\tprimes now has ", len(primes))
			fmt.Println("\tlargest prime: ", primes[len(primes) - 1])
		}
		return cur
	}
}()

func maxPrimeFactor(n *big.Int) *big.Int {
	fmt.Println("Working on ", n)
	for {
		i := nextPrime()
		for bigInt.Rem(n, i).Cmp(big.NewInt(0)) == 0 {
			if n.Cmp(i) == 0 { return i }
			fmt.Println("divisible by ", i)
			n.Div(n, i)
			fmt.Println("\t**\tn = ", n)
		}
	}

	panic("Unreachable"); // Please the compiler
}

func main() {
	//fmt.Println(maxPrimeFactor(big.NewInt(13195)))
	//fmt.Println(maxPrimeFactor(big.NewInt(600851475143)))

	i := new(big.Int)
	_, err := fmt.Sscan("600851475143601", i)
	if err != nil {
		fmt.Println("error scanning value:", err)
	}
	t = time.Now()
	fmt.Println(maxPrimeFactor(i))
}
