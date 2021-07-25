package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		if root.Left == nil {
			te := root.Right
			root = nil
			return te
		} else if root.Right == nil {
			te := root.Left
			root = nil
			return te
		}
		te := mv(root.Right)
		root.Data = te.Data
		root.Right = BTreeDeleteNode(root.Right, te)
	}
	return root
}

func mv(node *TreeNode) *TreeNode {
	c := node
	for c != nil && c.Left != nil {
		c = c.Left
	}
	return c
}
