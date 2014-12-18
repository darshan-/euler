package main

func reverse(n int) (rev int) {
	for place := 1; n > 0; place *= 10 {
		remainder := n % (place * 10)
		digit := remainder / place
		rev *= 10
		rev += digit
		n -= remainder
	}

	return rev
}

func allOdd(n int) bool {
	for place := 1; n > 0; place *= 10 {
		remainder := n % (place * 10)
		digit := remainder / place
		if digit % 2 == 0 { return false }
		n -= remainder
	}

	return true
}

const GOAL = 1000 * 1000 * 1000

func main() {
	var count int64

	place := 10
	for i := 10; i < GOAL; i++ {
		if i == place * 10    { place = i }
		if i % 10 == 0        { continue  }
		if i % 10 > i / place { continue  }

		rev := reverse(i)
		if rev >= i { continue }

		if allOdd(i + rev) {
				count += 2
		}
	}

	println(count)
}

/*
             1 ->       0
            10 ->       0
           100 ->      20
         1,000 ->     120
        10,000 ->     720
       100,000 ->     720
     1,000,000 ->  18,720
    10,000,000 ->  68,720
   100,000,000 -> 608,720
 1,000,000,000 ->(608,720)?? Maybe? -> Yes! Lucky guess!
*/
