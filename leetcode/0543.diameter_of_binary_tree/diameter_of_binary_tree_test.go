package leetcode

import (
	"crackleetgo/gods/tree"
	"testing"
)

func Test_diameterOfBinaryTree(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"1 2 4 x x 5 x x 3 x x", 3},
	}

	for _, e := range tests {
		result := diameterOfBinaryTree(tree.BuildTree(e.input))
		if result != e.expected {
			t.Errorf(
				"The diameter of binary tree(%s) is %d, but got %d",
				e.input,
				e.expected,
				result,
			)
		}
	}
}
