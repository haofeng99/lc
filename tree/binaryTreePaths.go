package tree

import "strconv"

// https://leetcode.cn/problems/binary-tree-paths/
// 257. 二叉树的所有路径
// string 值传递
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{"[]"}
	}

	ans := []string{}

	var dfs func(*TreeNode, string)
	dfs = func(node *TreeNode, s string) {
		if node == nil {
			return
		}
		s += ("->" + strconv.Itoa(node.Val))
		if node.Left == nil && node.Right == nil {
			if s[0:2] == "->" {
				s = s[2:]
			}
			ans = append(ans, s)
			return
		}
		dfs(node.Left, s)
		dfs(node.Right, s)
	}

	dfs(root, "")

	return ans
}
