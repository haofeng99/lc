package tree

// https://leetcode.cn/problems/minimum-absolute-difference-in-bst/
// 530. 二叉搜索树的最小绝对差
func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}

	abs := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}

	var pre *TreeNode
	ans := 1000000001
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre == nil {
			pre = node
		} else {
			ans = min(abs(node.Val-pre.Val), ans)
			pre = node
		}
		dfs(node.Right)
	}

	dfs(root)
	return ans
}
