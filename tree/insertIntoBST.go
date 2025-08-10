package tree

// https://leetcode.cn/problems/insert-into-a-binary-search-tree/
// 701.二叉搜索树中的插入操作
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	pre, cur := root, root
	for cur != nil {
		pre = cur
		if cur.Val > val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	if pre.Val < val {
		pre.Right = &TreeNode{Val: val}
	} else {
		pre.Left = &TreeNode{Val: val}
	}
	return root
}

func insertIntoBST_dfs(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val > val {
		root.Left = insertIntoBST_dfs(root.Left, val)
	} else {
		root.Right = insertIntoBST_dfs(root.Right, val)
	}
	return root
}
