package main

import (
	"fmt"
	"leetcode/array"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	s := array.Shuffle(arr, 3)
	fmt.Print(s)
}
