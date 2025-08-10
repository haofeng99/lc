package tree

// 145. 二叉树的后序遍历
// https://leetcode.cn/problems/binary-tree-postorder-traversal/description/
func postorderTraversal(root *TreeNode) []int {
	ans := []int{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		ans = append(ans, node.Val)
	}

	dfs(root)
	return ans
}

// 前序遍历 root -> 左 -> 右
// 中序遍历 左 -> root -> 右
// 后序遍历 左 -> 右 -> root
func postorderTraversal_bfs(root *TreeNode) []int {
	ans := []int{}
	queue := []*TreeNode{}
	for len(queue) > 0 || root != nil {
		for root != nil {
			ans = append([]int{root.Val}, ans...)
			queue = append([]*TreeNode{root}, queue...)
			root = root.Right
		}
		root = queue[0]
		root = root.Left
		queue = queue[1:]
	}
	return ans
}
