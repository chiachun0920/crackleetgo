package leetcode

import "testing"

func Test_twoSum(t *testing.T) {
	tests := []struct {
		numbers  []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
	}

	for _, e := range tests {
		result := twoSum(e.numbers, e.target)
		expectedResult := e.expected

		if !isSliceEqual(result, expectedResult) {
			t.Errorf("Given array of %v and target %d, ", e.numbers, e.target)
			t.Errorf(
				"and expected found index is (%d, %d) ",
				expectedResult[0],
				expectedResult[1],
			)
			t.Errorf(
				"but got index (%d, %d)",
				result[0],
				result[1],
			)
		}

	}
}

func isSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, value := range a {
		if value != b[idx] {
			return false
		}
	}

	return true
}
