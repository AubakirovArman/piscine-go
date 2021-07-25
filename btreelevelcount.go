package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root != nil {
		if BTreeLevelCount(root.Left) > BTreeLevelCount(root.Right) {
			return BTreeLevelCount(root.Left) + 1
		} else {
			return BTreeLevelCount(root.Right) + 1
		}
	} else {
		return 0
	}
}
