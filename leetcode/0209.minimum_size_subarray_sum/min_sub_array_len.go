package leetcode

import "math"

func minSubArrayLen(target int, nums []int) int {
	windowStart := 0

	sum := 0
	minLen := math.MaxInt

	for windowEnd, num := range nums {
		sum = sum + num
		for sum >= target {
			minLen = min(windowEnd-windowStart+1, minLen)
			// shrink the window size
			sum = sum - nums[windowStart]
			windowStart = windowStart + 1
		}
	}

	if minLen == math.MaxInt {
		return 0
	}

	return minLen
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
