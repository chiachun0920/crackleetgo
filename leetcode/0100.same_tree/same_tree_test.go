package leetcode

import (
	"crackleetgo/gods/tree"
	"testing"
)

func Test_isSameTree(t *testing.T) {
	var tests = []struct {
		inputX   string
		inputY   string
		expected bool
	}{
		{"1 2 x x 3 x x", "1 2 x x 3 x x", true},
		{"1 2 x x x x", "1 x x 2 x x", false},
		{"1 2 x x 1 x x", "1 1 x x 2 x x", false},
	}

	for _, e := range tests {
		result := isSameTree(tree.BuildTree(e.inputX), tree.BuildTree(e.inputY))
		if result != e.expected {
			t.Fatalf(
				"Tree %s and %s expected the same: %t, but got %t",
				e.inputX,
				e.inputY,
				result,
				e.expected,
			)
		}

	}
}
