package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root.Left == nil || root.Right == nil {
		return root
	}
	if BTreeMax(root.Left).Data > BTreeMax(root.Right).Data {
		return BTreeMax(root.Left)
	} else {
		return BTreeMax(root.Right)
	}
}
