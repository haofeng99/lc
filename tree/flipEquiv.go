package tree

// https://leetcode.cn/problems/flip-equivalent-binary-trees/
// 951 翻转等价二叉树
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil && root2 != nil {
		return false
	}
	if root1 != nil && root2 == nil {
		return false
	}

	// 1.值相等
	// 2.[翻转或不翻转]
	return root1.Val == root2.Val &&
		((flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left)) ||
			(flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)))
}
