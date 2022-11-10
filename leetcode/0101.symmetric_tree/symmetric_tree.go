package leetcode

import (
	"crackleetgo/gods/tree"
)

type TreeNode = tree.TreeNode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queueA := []*TreeNode{root}
	queueB := []*TreeNode{root}

	for len(queueA) > 0 {
		// dequeue both
		nodeA, nodeB := queueA[0], queueB[0]
		queueA, queueB = queueA[1:], queueB[1:]

		equal, isNil := nodeEqual(nodeA.Left, nodeB.Right)
		if !equal {
			return false
		}
		if !isNil {
			queueA = append(queueA, nodeA.Left)
			queueB = append(queueB, nodeB.Right)
		}

		equal, isNil = nodeEqual(nodeA.Right, nodeB.Left)
		if !equal {
			return false
		}
		if !isNil {
			queueA = append(queueA, nodeA.Right)
			queueB = append(queueB, nodeB.Left)
		}
	}

	if len(queueB) > 0 {
		return false
	}

	return true
}

func nodeEqual(nodeA *TreeNode, nodeB *TreeNode) (bool, bool) {
	if nodeA == nil && nodeB == nil {
		return true, true
	}

	if nodeA != nil && nodeB != nil && nodeA.Val == nodeB.Val {
		return true, false
	}

	return false, false
}
