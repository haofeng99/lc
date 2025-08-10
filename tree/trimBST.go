package tree

// https://leetcode.cn/problems/trim-a-binary-search-tree/
// 669. 修剪二叉搜索树

// dfs O(n) O(n)
func trimBST(root *TreeNode, low int, high int) *TreeNode {

	var dfs func(*TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		if node.Val < low {
			return dfs(node.Right)
		}
		if node.Val > high {
			return dfs(node.Left)
		}
		node.Left = dfs(node.Left)
		node.Right = dfs(node.Right)
		return node
	}

	return dfs(root)
}

// bfs O(n) O(1)
func trimBST_bfs(root *TreeNode, low int, high int) *TreeNode {
	// 1.先定root
	for root != nil && (root.Val < low || root.Val > high) {
		if root.Val < low {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	if root == nil {
		return nil
	}

	// 左子树修剪
	node := root
	for node != nil {
		if node.Left != nil && node.Left.Val < low {
			node.Left = node.Left.Right // 真正的剪切
		} else {
			node = node.Left
		}
	}

	// 右子树修剪
	node = root
	for node != nil {
		if node.Right != nil && node.Right.Val > high {
			node.Right = node.Right.Left // 真正的剪切
		} else {
			node = node.Right
		}
	}

	return root
}
