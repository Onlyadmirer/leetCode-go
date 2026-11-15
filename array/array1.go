package array

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
