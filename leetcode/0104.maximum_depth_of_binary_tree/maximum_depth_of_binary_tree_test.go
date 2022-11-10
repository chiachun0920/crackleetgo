package leetcode

import (
	"crackleetgo/gods/tree"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"3 9 x x 20 15 x x 7 x x", 3},
		{"1 x 2 x x", 2},
	}

	for _, e := range tests {
		result := maxDepth(tree.BuildTree(e.input))
		if result != e.expected {
			t.Errorf(
				"Expect depth of tree %s is %d but got %d",
				e.input,
				e.expected,
				result,
			)
		}
	}
}
