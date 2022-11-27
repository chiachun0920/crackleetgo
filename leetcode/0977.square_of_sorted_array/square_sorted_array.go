package leetcode

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))

	cursor := len(nums) - 1
	left, right := 0, len(nums)-1

	for left <= right {
		leftVal, rightVal := nums[left], nums[right]
		if leftVal*leftVal > rightVal*rightVal {
			result[cursor] = leftVal * leftVal
			left++
		} else {
			result[cursor] = rightVal * rightVal
			right--

		}
		cursor--
	}
	return result
}
