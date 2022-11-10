package leetcode

import (
	"crackleetgo/gods/tree"
	"testing"
)

func Test_isBalanced(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"3 9 x x 20 15 x x 7 x x", true},
		{"1 2 3 4 x x 4 x x 3 x x 2 x x", false},
		{"1 x 2 x 3 x x", false},
		{"x", true},
	}
	for _, e := range tests {
		result := isBalanced(tree.BuildTree(e.input))
		if result != e.expected {
			t.Errorf(
				"Expect tree %s is balanced: %t, but got %t",
				e.input,
				e.expected,
				result,
			)
		}
	}
}
