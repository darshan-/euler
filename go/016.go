package main

import "math/big"

func main() {
	s := new(big.Int).Exp(big.NewInt(2), big.NewInt(1000), nil).String();
	sum := 0
	for _, b := range []byte(s) {
		sum += int(b - '0')
	}
	println(sum)
}
