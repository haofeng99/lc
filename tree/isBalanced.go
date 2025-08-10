package tree

// https://leetcode.cn/problems/balanced-binary-tree/
// 最简单的思路, 自顶向下遍历, 使用求节点最大深度方式求解
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if abs(maxDepth(root.Left)-maxDepth(root.Right)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}
func abs(v int) int {
	if v <= 0 {
		return -v
	}
	return v
}

// 后续遍历 时空复杂度O(n)
func isBalanced_post_order(root *TreeNode) bool {
	return heigh(root) > 0
}

func heigh(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftHigh := heigh(node.Left)
	rightHigh := heigh(node.Right)
	if leftHigh == -1 || rightHigh == -1 || abs(leftHigh-rightHigh) > 1 {
		return -1
	}
	return max(leftHigh, rightHigh) + 1
}
