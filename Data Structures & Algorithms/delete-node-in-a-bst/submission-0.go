/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findMinNode(root *TreeNode) *TreeNode {
	curr := root

	for curr != nil && curr.Left != nil {
		curr = curr.Left
	}
	return curr
}
 
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
		return nil // no child
	}

	// find node to delete
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else { // node found
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else { // 2 children
			minNode := findMinNode(root.Right)
			root.Val = minNode.Val
			root.Right = deleteNode(root.Right, minNode.Val)
		}
	}
	return root
}
