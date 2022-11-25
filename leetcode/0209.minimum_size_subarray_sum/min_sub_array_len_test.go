package leetcode

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	tests := []struct {
		target int
		nums   []int
		expect int
	}{
		{7, []int{2, 3, 1, 2, 4, 3}, 2},
		{4, []int{1, 4, 4}, 1},
		{11, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 0},
	}

	for _, e := range tests {
		if e.expect != minSubArrayLen(e.target, e.nums) {
			t.Errorf("Answer of minSubArrayLen not passed")
		}
	}
}
