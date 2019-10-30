package exercises

import (
	"fmt"
)

//TreeNode https://leetcode-cn.com/problems/sum-of-left-leaves/
type TreeNode struct {
	val       int
	leftNode  *TreeNode
	rightNode *TreeNode
}

// RootNode 根节点
func RootNode(val int) *TreeNode {
	return &TreeNode{val, nil, nil}
}

func (node *TreeNode) addLeftNode(leftNode *TreeNode) *TreeNode {
	node.leftNode = leftNode
	return leftNode
}
func (node *TreeNode) addRightNode(rightNode *TreeNode) *TreeNode {
	node.rightNode = rightNode
	return rightNode
}
func (node TreeNode) String() string {
	return fmt.Sprintf("%#v [left]=%#v[right]=%#v", node.val, node.leftNode.val, node.rightNode.val)
}
func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	sumLeft(root, &sum)
	return sum
}

//条件结束的为子节点为nil
func sumLeft(node *TreeNode, sum *int) {
	if node == nil {
		return
	}
	sumLeft(node.rightNode, sum)

	if node.leftNode != nil {
		*sum += node.leftNode.val
	}
	sumLeft(node.leftNode, sum)
}
