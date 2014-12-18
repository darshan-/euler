package main

const GOAL = 100 * 1000
const SQRT = 317

var primes []int
var lp int
var precalc [GOAL + 1]int

func genPrimes() {
	var numbers [GOAL + 1]bool

	for i := 2; i <= SQRT; {
		primes = append(primes, i)

		for j := i * i; j <= GOAL; j += i {
			numbers[j] = true
		}

		for i++; numbers[i]; i++ {}
	}

	for i := SQRT + 1; i <= GOAL; i++ {
		if ! numbers[i] {
			primes = append(primes, i)
		}
	}

	lp = len(primes)
}

var factors [20]int
func rad(n int) int {
	r := 1
	m := n
	l := 0

	for i := 0; i < lp; i++ {
		p := primes[i]

		if m % p == 0 {
			if l < 20 {
				factors[l] = p
				l++
			}
			r *= p
			for m % p == 0 {
				if m == p { goto done }
				m /= p
			}
		}
	}

	done:
	for i, f := range factors {
		if f == 0 { break }
		for q := uint64(n * f); q <= GOAL ; q *= uint64(f) {
			precalc[q] = r
			factors[i] = 0
		}
	}

	return r
}

func E(n int, rads map[int] []int) int {
	count := 0
	var i int
	for i = 1; count < n; i++ {
		count += len(rads[i])
	}
	i--
	count -= len(rads[i])
	count++

	return rads[i][n - count]
}

func main() {
	rads := map[int] []int {}
	genPrimes()

	rads[1] = append(rads[1], 1)
	for i := 2; i <= GOAL; i++ {
		var r int
		if precalc[i] > 0 {
			r = precalc[i]
		} else {
			r = rad(i)
		}
		rads[r] = append(rads[r], i)
	}

	println(E(10000, rads))
}
