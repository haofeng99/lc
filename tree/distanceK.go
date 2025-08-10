package tree

// https://leetcode.cn/problems/all-nodes-distance-k-in-binary-tree/
// 863. 二叉树中所有距离为 K 的结点
func distanceK(root *TreeNode, target *TreeNode, k int) []int {

	ans := []int{}

	var dfs func(*TreeNode, []int)
	dfs = func(node *TreeNode, path []int) {
		if node == nil {
			return
		}
		if node.Val == target.Val {
			if len(path) >= k {
				ans = append(ans, path[len(path)-k])
			}
		} else {
			if len(path) >= k {
				if path[len(path)-k-1] == target.Val {
					ans = append(ans, node.Val)
				}
			}
		}

		path = append(path, node.Val)
		dfs(node.Left, path)
		dfs(node.Right, path)

		path = path[:len(path)-1]
	}

	dfs(root, []int{})

	return ans
}
