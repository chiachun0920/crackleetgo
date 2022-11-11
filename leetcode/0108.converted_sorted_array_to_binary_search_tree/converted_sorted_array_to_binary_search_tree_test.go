package leetcode

import "testing"

func Test_sortedArrayToBST(t *testing.T) {
	var tests = []struct {
		input    []int
		expected string
	}{
		{[]int{-10, -3, 0, 5, 9}, "0 -10 x -3 x x 5 x 9 x x"},
	}

	for _, e := range tests {
		node := sortedArrayToBST(e.input)
		encodeTree := node.Encode()
		if encodeTree != e.expected {
			t.Errorf(
				"Expect highest balanced tree is %s, but got %s",
				e.expected,
				encodeTree,
			)
		}
	}
}
