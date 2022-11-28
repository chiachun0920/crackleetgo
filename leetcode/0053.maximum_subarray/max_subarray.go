package leetcode

import "math"

func maxSubArray(nums []int) int {
	maxSubArraySum := math.MinInt
	// k: window size
	for k := 1; k <= len(nums); k++ {
		kMaxSum := kMaxSubArray(nums, k)
		maxSubArraySum = max(maxSubArraySum, kMaxSum)
	}
	return maxSubArraySum
}

func kMaxSubArray(nums []int, k int) int {
	sum := 0
	windowStart := 0
	maxSum := math.MinInt
	for idx, value := range nums {
		sum += value
		if idx >= k-1 {
			// check if sum is max
			maxSum = max(sum, maxSum)
			// shrink by 1
			sum -= nums[windowStart]
			windowStart++
		}
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
