package leetcode

import "crackleetgo/gods/tree"

type TreeNode = tree.TreeNode

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// switch left and right
	left := root.Left
	root.Left = root.Right
	root.Right = left

	root.Left = invertTree(root.Left)
	root.Right = invertTree(root.Right)
	return root
}
