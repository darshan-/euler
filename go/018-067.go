package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

const filename = "../data/067-triangle.txt"

func readFile() (nums [][]int) {
	f, err := os.Open(filename)
	if err != nil { panic(err) }
	defer f.Close()

	rd := bufio.NewReader(f)

	for {
		line, isPrefix, err := rd.ReadLine()
		if isPrefix      { panic("Woops, isPrefix is true!") }
		if err == io.EOF { break }
		if err != nil    { panic(err) }

		buf := bytes.NewBuffer(line)

		var lineNums []int

		for {
			var i int
			_, err := fmt.Fscanf(buf, "%d", &i)
			if err != nil { break }
			lineNums = append(lineNums, i)
		}

		nums = append(nums, lineNums)
	}

	return nums
}

func max(a, b int) int {
	if a > b { return a }
	return b
}

func main() {
	nums := readFile()

	for r := len(nums) - 2; r >= 0; r-- {
		for c := 0; c < len(nums[r]); c++ {
			nums[r][c] += max(nums[r+1][c], nums[r+1][c+1])
		}
	}

	fmt.Println(nums[0][0])
}
