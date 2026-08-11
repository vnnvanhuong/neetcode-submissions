/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    stack := []*TreeNode{}
	cur := root
	cnt := 0
	
	for len(stack) > 0 || cur != nil {
		// push left
		for cur != nil {
			stack=append(stack, cur)
			cur = cur.Left
		}

		// process the check
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cnt++

		if cnt == k {
			return cur.Val
		}

		// push right
		cur = cur.Right
	}

	return 0
}
