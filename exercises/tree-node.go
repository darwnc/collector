package exercises

import (
	"fmt"
	"strconv"
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

// https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/
func preorder(node *TreeNode) {
	// nodes := []*TreeNode{node}
	var leftNode []*TreeNode
	var rightNode []*TreeNode
	// collect.
	if node.leftNode != nil {
		leftNode = []*TreeNode{node.leftNode}
	}
	if node.rightNode != nil {
		rightNode = []*TreeNode{node.rightNode}
	}
	left := []string{}
	right := []string{}
	for len(leftNode) != 0 || len(rightNode) != 0 {
		if len(leftNode) != 0 {
			if leftNode[0].leftNode != nil {
				leftNode = append(leftNode, leftNode[0].leftNode)
			}

			if leftNode[0].rightNode != nil {
				rightNode = append(rightNode, leftNode[0].rightNode)
			}
			left = append(left, strconv.Itoa(leftNode[0].val))
			// fmt.Println("left=", leftNode[0].val)
			leftNode = leftNode[1:]
		}

		if len(rightNode) != 0 {
			if rightNode[0].leftNode != nil {
				leftNode = append(leftNode, rightNode[0].leftNode)
			}

			if rightNode[0].rightNode != nil {
				rightNode = append(rightNode, rightNode[0].rightNode)
			}
			right = append(right, strconv.Itoa(rightNode[0].val))
			// fmt.Println("right=", rightNode[0].val)
			rightNode = rightNode[1:]
		}

	}
	fmt.Println(left)
	fmt.Println(right)
}
