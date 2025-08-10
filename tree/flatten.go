package tree

// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/
// 右子树的父节点为左子树的最右节点的右子节点
func flatten(root *TreeNode) {
	for root != nil {
		if root.Left != nil {
			subRoot := root.Left
			for subRoot.Right != nil {
				subRoot = subRoot.Right
			}
			subRoot.Right = root.Right
			root.Right = root.Left
			root.Left = nil
		}
		root = root.Right
	}
}
