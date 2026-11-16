package main

import (
	"fmt"
	"leetcode/array"
)

func main() {
	arr := []int{}
	arr2 := []int{2, 3}
	s := array.FindMedianSortedArrays(arr, arr2)
	fmt.Printf("%.2f\n", s)
}
