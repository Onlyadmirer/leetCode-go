package main

import (
	"fmt"
	"leetcode/array"
)

func main() {
	arr := []string{"flower", "flow", "flight"}
	s := array.LongestCommonPrefix(arr)
	fmt.Print(s)
}
