package tree

// https://leetcode.cn/problems/binary-tree-preorder-traversal/
func preorderTraversal(root *TreeNode) []int {
	return bfs_preorderTraversal(root)
}

func dfs_preorderTraversal(node *TreeNode, arr *[]int) {
	if node == nil {
		return
	}
	*arr = append(*arr, node.Val)
	dfs_preorderTraversal(node.Left, arr)
	dfs_preorderTraversal(node.Right, arr)
}

func bfs_preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	nodeList := []*TreeNode{}
	ans := []int{}
	for root != nil || len(nodeList) > 0 {
		for root != nil {
			nodeList = append(nodeList, root)
			ans = append(ans, root.Val)
			root = root.Left
		}
		root = nodeList[len(nodeList)-1]
		nodeList = nodeList[0 : len(nodeList)-1]
		root = root.Right
	}
	return ans
}
