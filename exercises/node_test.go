package exercises

import (
	"fmt"
	"testing"
)

func TestNode(t *testing.T) {
	rootNode := RootNode(3)
	leftNode := RootNode(10)
	rightNode := RootNode(14)
	rightNodeLeft := RootNode(12)
	rightNodeRight := RootNode(13)

	rightNodeRightLeft := RootNode(15)
	rightNodeLeft.addLeftNode(rightNodeRightLeft)

	rightNode.addLeftNode(rightNodeLeft)
	rightNode.addRightNode(rightNodeRight)
	rootNode.addRightNode(rightNode)
	rootNode.addLeftNode(leftNode)
	fmt.Println(rootNode)
	fmt.Println("sum=", sumOfLeftLeaves(rootNode))

	preorder(rootNode)
}
