package tree

// https://leetcode.cn/problems/maximum-depth-of-n-ary-tree/
// 559. N 叉树的最大深度
func maxDepthNSub(root *NodeNSub) int {
	if root == nil {
		return 0
	}
	ans := 0
	queue := []*NodeNSub{root}
	for len(queue) > 0 {
		tempQ := queue
		queue = nil
		for _, node := range tempQ {
			queue = append(queue, node.Children...)
		}
		ans++
	}
	return ans
}

func maxDepthNSub_dfs(root *NodeNSub) int {
	if root == nil {
		return 0
	}
	curDepth := 0
	for _, node := range root.Children {
		deep := maxDepthNSub_dfs(node)
		if deep > curDepth {
			curDepth = deep
		}
	}
	return curDepth + 1
}
