package leetcode

import (
	"crackleetgo/gods/tree"
)

type TreeNode = tree.TreeNode

func diameterOfBinaryTree(root *TreeNode) int {
	var max = 0
	return traverseAndEval(root, &max)
}

func traverseAndEval(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	length := calDiameter(root)
	if length > *max {
		*max = length
	}

	traverseAndEval(root.Left, max)
	traverseAndEval(root.Right, max)
	return *max
}

func calDiameter(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return treeHeight(root.Left) + treeHeight(root.Right)
}

func treeHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + max(treeHeight(node.Left), treeHeight(node.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
