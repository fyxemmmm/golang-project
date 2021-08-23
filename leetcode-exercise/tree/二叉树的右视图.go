package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		q := []*TreeNode{}
		arr := []int{}
		for _, node := range queue {
			if node == nil {
				continue
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}

			arr = append(arr, node.Val)
		}
		res = append(res, arr[len(arr)-1])
		queue = q
	}

	return res
}