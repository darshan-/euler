package main

func main() {
	squares := map[int] int {}

	for i := 0; i < 1000; i++ {
		squares[i * i] = i
	}

	for a := 1; a < 1000; a++ {
		for b := a + 1; b < 1000; b++ {
			c := squares[a * a + b * b]
			if c > 0 && a + b + c == 1000 {
				println(a * b *c)
			}
		}
	}
}
