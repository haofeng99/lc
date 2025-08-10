package tree

// https://leetcode.cn/problems/lowest-common-ancestor-of-deepest-leaves/
// 865. 具有所有最深节点的最小子树
// 1123. 最深叶节点的最近公共祖先
// lcaDeepestLeaves
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	var dfs func(*TreeNode) (*TreeNode, int)
	dfs = func(node *TreeNode) (*TreeNode, int) {
		if node == nil {
			return nil, 0
		}
		left, d1 := dfs(node.Left)
		right, d2 := dfs(node.Right)
		if d1 == d2 {
			return node, d1 + 1
		}
		if d1 > d2 {
			return left, d1 + 1
		}
		return right, d2 + 1
	}

	node, _ := dfs(root)
	return node
}
