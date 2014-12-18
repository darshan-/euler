package main

type list []uint64

func (l *list) add(n uint64) {
	*l = append(*l, n)
}

func (l *list) total() (sum uint64) {
	for _, n := range *l {
		sum += n
	}
	return sum
}

const dimen = 20

func main() {
	last := list {1}

	for j := 1; j < dimen; j++ {
		var cur list
		for i, n := range last {
			var subtotal uint64
			if (i > 0) { subtotal = cur[i-1] }

			cur.add(subtotal + n)
		}
		cur.add(last.total() * 2)
		last = cur
	}

	println(last.total() * 2)
}
