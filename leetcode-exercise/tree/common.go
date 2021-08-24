package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return  a
	}
	return b
}




func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
