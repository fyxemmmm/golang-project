package tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := 0
	var dfs func(node *TreeNode, preSum int)
	dfs = func(node *TreeNode, preSum int) {
		if node == nil {
			return
		}

		num := preSum * 10 + node.Val
		if node.Left == nil && node.Right == nil {
			res += num
			return
		}

		dfs(node.Left, num)
		dfs(node.Right, num)
	}

	dfs(root, 0)
	return res
}