package leetcode

import (
	"crackleetgo/gods/tree"
)

type TreeNode = tree.TreeNode

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if intAbs(treeHeight(root.Left)-treeHeight(root.Right)) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func intAbs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}

func treeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(treeHeight(root.Left), treeHeight(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
