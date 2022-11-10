package leetcode

import (
	"crackleetgo/gods/tree"
	"testing"
)

func Test_isSymmetric(t *testing.T) {
	var tests = []struct {
		inputTree string
		expected  bool
	}{
		{"1 0 x x x", false},
		{"1 2 3 x x 4 x x 2 4 x x 3 x x", true},
		{"1 2 x 3 x x 2 x 3 x x", false},
	}

	for _, e := range tests {
		result := isSymmetric(tree.BuildTree(e.inputTree))
		if result != e.expected {
			if e.expected {
				t.Errorf(
					"Expected tree %s is symmetric, but got %t",
					e.inputTree,
					result,
				)
			} else {
				t.Errorf(
					"Expected tree %s is asymmetric, but got %t",
					e.inputTree,
					result,
				)
			}
		}
	}
}
