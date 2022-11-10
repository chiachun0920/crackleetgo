package leetcode

import (
	"crackleetgo/gods/tree"
	"testing"
)

func Test_invertTree(t *testing.T) {
	var tests = []struct {
		root     string
		expected string
	}{
		{
			"4 2 1 x x 3 x x 7 6 x x 9 x x",
			"4 7 9 x x 6 x x 2 3 x x 1 x x",
		},
	}

	for _, e := range tests {
		n := invertTree(tree.BuildTree(e.root))
		result := n.Encode()
		if result != e.expected {
			t.Errorf(
				"Expected tree %s after inverting will be %s, but got %s",
				e.root,
				e.expected,
				result,
			)
		}
	}
}
