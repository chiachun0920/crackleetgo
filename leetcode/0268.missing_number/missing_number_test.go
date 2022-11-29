package leetcode

import "testing"

func Test_missingNumber(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
		{[]int{3, 0, 1}, 2},
	}

	for _, e := range tests {
		result := missingNumber(e.nums)
		if e.expected != result {
			t.Errorf(
				"Number of missing is not %d, expected %d",
				result,
				e.expected,
			)
		}
	}
}
