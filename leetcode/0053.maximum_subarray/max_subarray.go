package leetcode

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	curSum := 0
	for _, value := range nums {
		if curSum < 0 {
			curSum = 0
		}
		curSum += value
		maxSum = max(curSum, maxSum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
