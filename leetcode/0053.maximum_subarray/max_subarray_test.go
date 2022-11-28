package leetcode

import "testing"

func Test_maxSubArray(t *testing.T) {
	tests := []struct {
		nums   []int
		maxSum int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{-2}, -2},
	}

	for _, e := range tests {
		result := maxSubArray(e.nums)
		if result != e.maxSum {
			t.Errorf("Max sub array is not %d, expect %d", result, e.maxSum)
		}
	}
}
