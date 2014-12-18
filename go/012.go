package main

func countFactors(n int) int {
	count := 0

	for i := 1; ; i++ {
		if n % i == 0 { // i is a factor
			count++
			if i * i == n {
				return count * 2 - 1
			} else if i * i > n {
					return (count - 1) * 2
			}
		}
		if i * i > n {
				return count * 2
		}
	}

	panic("Never happens.")
}

const GOAL = 500

func main() {
	triNum := 0;
	for i := 1; ; i++ {
		triNum += i
		if countFactors(triNum) > GOAL { break }
	}
	println(triNum)
}
