package leetcode

import (
	"fmt"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{
			[]int{-4, -1, 0, 3, 10},
			[]int{0, 1, 9, 16, 100},
		},
	}

	for _, e := range tests {
		result := sortedSquares(e.nums)
		if !isSliceEqual(result, e.expected) {
			messages := []string{
				"Answer of sortedSquares is not passed!",
				fmt.Sprintf("Expect get %v but got %v", e.expected, result),
			}
			t.Errorf("%s %s", messages[0], messages[1])
		}
	}
}

func isSliceEqual(a, b []int) bool {
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
