package tree

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Build a tree using string pattern like 1 2 x x 3 x x
func BuildTree(s string) *TreeNode {
	pos := 0
	return concatNode(splitWords(s), &pos)
}

// Use pre-order traversal to encode tree
func (n *TreeNode) Encode() string {
	if n == nil {
		return "x"
	}
	return fmt.Sprintf("%d %s %s", n.Val, n.Left.Encode(), n.Right.Encode())
}

func concatNode(nodes []string, pos *int) *TreeNode {
	val := nodes[*pos]
	*pos++
	if val == "x" {
		return nil
	}
	v, _ := strconv.Atoi(val)
	left := concatNode(nodes, pos)
	right := concatNode(nodes, pos)
	return &TreeNode{v, left, right}
}

func splitWords(s string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, " ")
}
