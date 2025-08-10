package tree

import "math"

// https://leetcode.cn/problems/validate-binary-search-tree/
// 98. 验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	var dfs func(*TreeNode, int, int) bool
	dfs = func(node *TreeNode, left, right int) bool {
		if node == nil {
			return true
		}
		if node.Val <= left || node.Val >= right {
			return false
		}
		return dfs(node.Left, left, node.Val) && dfs(node.Right, node.Val, right)
	}
	return dfs(root, math.MinInt64, math.MaxInt64)
}

// 二叉搜索树-中序遍历一定是有序的

func inorder_isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	nodeList := []*TreeNode{}
	curNodeVal := math.MinInt64
	for root != nil || len(nodeList) > 0 {
		for root != nil {
			nodeList = append(nodeList, root)
			root = root.Left
		}
		root = nodeList[len(nodeList)-1]
		nodeList = nodeList[:len(nodeList)-1]
		if root.Val <= curNodeVal {
			return false
		}
		curNodeVal = root.Val

		root = root.Right

	}
	return true
}
