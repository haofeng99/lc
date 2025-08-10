package tree

// https://leetcode.cn/problems/delete-leaves-with-a-given-value/
// 1325. 删除给定值的叶子节点
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)
	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	} else {
		return root
	}
}
