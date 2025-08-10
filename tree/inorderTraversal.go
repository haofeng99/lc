package tree

// https://leetcode.cn/problems/binary-tree-inorder-traversal/
func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	dfs_inorderTraversal(root, &ans)
	return ans
}

func dfs_inorderTraversal(node *TreeNode, ans *[]int) {
	if node == nil {
		return
	}
	if node.Left != nil {
		dfs_inorderTraversal(node.Left, ans)
	}
	*ans = append(*ans, node.Val)
	if node.Right != nil {
		dfs_inorderTraversal(node.Right, ans)
	}
}

// 滴滴之痛
func bfs_inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{}
	ans := []int{}
	for root != nil || len(queue) > 0 {
		for root != nil {
			queue = append(queue, root)
			root = root.Left
		}
		root = queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		ans = append(ans, root.Val)
		root = root.Right
	}
	return ans
}
