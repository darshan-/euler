package main

const BOUND = 1000 * 1000

func main() {
	var max, which int
	for i := 1; i < BOUND; i++ {
		count := 1
		for j := uint64(i); j > 1; count++ {
			if j % 2 == 0 {
				j = j / 2
			} else {
				j = j * 3 + 1
			}
		}
		if count > max {
			max = count
			which = i
			//println("Max:", max, "starts at", which)
		}
	}
	println("Max:", max, "starts at", which)
}
