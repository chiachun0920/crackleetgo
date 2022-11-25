package leetcode

import "testing"

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}

	for _, e := range tests {
		expectedResult := e.output
		actualResult := removeDuplicates(e.input)

		if expectedResult != actualResult {
			t.Errorf(
				"Expected output: %d, but got %d",
				expectedResult,
				actualResult,
			)
		}
	}
}
