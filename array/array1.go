package array

import (
	"sort"
)

func GetConcatenation(nums []int) []int {

	n := len(nums)
	ans := make([]int, n*2)

	for i := range nums {
		ans[i] = nums[i]
		ans[i+n] = nums[i]
	}

	return ans
}

func Shuffle(nums []int, n int) []int {

	ans := make([]int, len(nums)/2)
	arr := make([]int, 0, len(nums))

	for i := range ans {
		arr = append(arr, nums[i])
		arr = append(arr, nums[i+n])
	}

	return arr
}

// func TwoSum(nums []int, target int) []int {

// 	var res []int

// 	for i, val := range nums {
// 		for j, val2 := range nums[i+1:] {
// 			if val+val2 == target {
// 				res = append(res, i, i+j+1)
// 				break
// 			}

// 		}

// 	}

// 	return res
// }

func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if val, found := numMap[complement]; found {
			return []int{val, i}
		}
		numMap[num] = i
	}
	return nil
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr := append(nums1, nums2...)

	if len(arr) == 1 {
		return float64(arr[0])
	}

	sort.Ints(arr)

	var res float64

	res2 := (len(arr) + 1) / 2
	if len(arr)%2 != 0 {
		res = float64(arr[res2-1])
	} else if arr[(len(arr)/2)] == 0 && arr[len(arr)/2-1] == 0 {
		res = 0
	} else {
		res = float64(arr[(len(arr)/2)]+arr[len(arr)/2-1]) / 2
	}

	return res

}
