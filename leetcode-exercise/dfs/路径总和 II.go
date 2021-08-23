package dfs
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}

	var dfs func(node *TreeNode, path []int, presum int)
	dfs = func(node *TreeNode, path []int, presum int) {
		if node == nil {
			return
		}

		path = append(path, node.Val)
		num := presum + node.Val

		if node.Left == nil && node.Right == nil {
			if num == targetSum {
				res = append(res, append([]int{}, path...))
				return
			}
		}

		dfs(node.Left, path, num)
		dfs(node.Right, path, num)

	}

	dfs(root, []int{}, 0)

	return res
}