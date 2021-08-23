package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	var dfs func(node *TreeNode, sum int) bool

	dfs = func(node *TreeNode, presum int) bool {
		if node == nil {
			return false
		}

		sum := presum + node.Val
		if node.Left == nil && node.Right == nil {
			if sum == targetSum {
				return true
			}else {
				return false
			}
		}

		return dfs(node.Left, sum) || dfs(node.Right, sum)

	}

	return dfs(root, 0)
}