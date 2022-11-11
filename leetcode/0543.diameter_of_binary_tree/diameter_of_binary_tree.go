package leetcode

import (
	"crackleetgo/gods/tree"
)

type TreeNode = tree.TreeNode

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lHeight := treeHeight(root.Left)
	rHeight := treeHeight(root.Right)

	diameterNode := lHeight + rHeight + 2

	return max(
		diameterNode,
		max(diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right)),
	)
}

func treeHeight(root *TreeNode) int {
	if root == nil {
		return -1
	}
	return 1 + max(treeHeight(root.Left), treeHeight(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
