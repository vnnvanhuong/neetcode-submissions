/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    res := []int{}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		qlen := len(queue)
		var rightMost *TreeNode
		for i := 0; i < qlen; i++ {
			node := queue[0]
			queue = queue[1:]

			if node != nil {
				rightMost = node
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			}

		}

		if rightMost != nil {
			res = append(res, rightMost.Val)
		}
	}

	return res
}
