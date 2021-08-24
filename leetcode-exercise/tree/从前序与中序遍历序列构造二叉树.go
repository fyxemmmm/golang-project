package tree


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := preorder[0]
	rootIndex := 0
	for idx, value := range inorder {
		if root == value {
			rootIndex = idx
			break
		}
	}

	preLeft, preRight := preorder[1:rootIndex+1], preorder[rootIndex+1:]
	inLeft, inRight := inorder[0:rootIndex], inorder[rootIndex+1:]
	return &TreeNode{
		Val: root,
		Left: buildTree(preLeft, inLeft),
		Right: buildTree(preRight, inRight),
	}
}