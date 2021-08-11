package bfs

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	queue := []*TreeNode{root}

	for i:= 0; len(queue) > 0; i ++ {
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

		if i % 2 == 1 {
			length := len(arr)
			for i := 0; i < length/2; i ++ {
				arr[i], arr[length-i-1] = arr[length-i-1], arr[i]
			}
		}

		queue = q
		res = append(res, arr)
	}

	return res
}