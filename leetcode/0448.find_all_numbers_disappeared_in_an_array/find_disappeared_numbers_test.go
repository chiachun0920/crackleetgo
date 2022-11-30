package leetcode

import "testing"

func Test_findDisappearedNumbers(t *testing.T) {
	tests := []struct {
		nums   []int
		output []int
	}{
		{[]int{2, 3, 1, 8, 2, 3, 5, 1}, []int{4, 6, 7}},
	}

	for _, e := range tests {
		result := findDisappearedNumbers(e.nums)
		if !sliceEqual(e.output, result) {
			t.Errorf(
				"Answer of find disappeared wrong, expect %v but got %v",
				e.output,
				result,
			)
		}
	}
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, value := range a {
		if b[idx] != value {
			return false
		}
	}

	return true
}
