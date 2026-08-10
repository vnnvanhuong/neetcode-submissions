/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    b, _ := dfs(root)
	return b
}

func dfs(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	lb, lh := dfs(root.Left)
	rb, rh := dfs(root.Right)

	b := lb && rb && abs(lh - rh) <= 1
	h := 1 + max(lh, rh)
	return b, h
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}