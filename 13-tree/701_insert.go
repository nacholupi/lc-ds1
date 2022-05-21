package tree

/* 701. Insert into a Binary Search Tree */
func insertIntoBST(root *TreeNode, val int) *TreeNode {

	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}
