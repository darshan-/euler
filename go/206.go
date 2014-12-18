package main

import "runtime"

const START = 1010101010
const END   = 1389026620

var check = map[uint64] uint64 {
	1000000000000000000: 1,
	10000000000000000: 2,
	100000000000000: 3,
	1000000000000: 4,
	10000000000: 5,
	100000000: 6,
	1000000: 7,
	10000: 8,
	100 : 9,
	1: 0,
}

func match(n uint64) bool {
	for place := uint64(1); n > 0; place *= 100 {
		remainder := n % (place * 10)
		digit := remainder / place
		if digit != check[place] { return false }
		n -= remainder
	}

	return true
}

func search(id, total int, c chan uint64) {
	var i uint64
	for i = uint64(START + id * 10); i < END; i += uint64(total * 10) {
		if match(i * i) {
			c <- i
			return
		}
	}
}

func main() {
	nProcs := runtime.NumCPU()
	runtime.GOMAXPROCS(nProcs)
	c := make(chan uint64)

	for i := 0; i < nProcs; i++ {
		go search(i, nProcs, c)
	}

	println(<-c)
}
