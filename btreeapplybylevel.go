package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	i := 0
	for bTreeApplyByLevel(root, i, f) {
		i++
	}
}

func bTreeApplyByLevel(root *TreeNode, level int, f func(...interface{}) (int, error)) bool {
	if root == nil {
		return false
	}
	if level == 0 {
		f(root.Data)
		return true
	}
	left := bTreeApplyByLevel(root.Left, level-1, f)
	right := bTreeApplyByLevel(root.Right, level-1, f)
	return left || right
}
