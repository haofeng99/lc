package tree

// https://leetcode.cn/problems/find-mode-in-binary-search-tree/
// 501. 二叉搜索树中的众数
// 中序遍历有序
func findMode(root *TreeNode) []int {
	curCnt, curVal, maxCnt := 0, -100000001, 0
	ans := []int{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if node.Val == curVal {
			curCnt++
		} else {
			curCnt = 1
			curVal = node.Val
		}
		if curCnt == maxCnt {
			ans = append(ans, node.Val)
		}
		if curCnt > maxCnt {
			maxCnt = curCnt
			ans = []int{node.Val}
		}
		dfs(node.Right)
	}

	dfs(root)

	return ans
}
