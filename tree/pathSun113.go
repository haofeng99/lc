package tree

// https://leetcode.cn/problems/path-sum-ii/
// dfs + bfs 都要会
// 113. 路径总和 II
func pathSum113(root *TreeNode, target int) [][]int {
	ans := [][]int{}
	var dfs func(*TreeNode, []int, int)
	dfs = func(node *TreeNode, curPath []int, curSum int) {
		if node == nil {
			return
		}
		curPath = append(curPath, node.Val)
		curSum = curSum + node.Val

		if node.Left == nil && node.Right == nil && curSum == target {
			// 在 ans = append(ans, curPath) 这行，如果直接把 curPath 加入了答案，
			// 但 curPath 是复用的切片，后续递归会修改它，导致 ans 里的路径被后续递归污染。
			pathCopy := make([]int, len(curPath))
			copy(pathCopy, curPath)
			ans = append(ans, pathCopy)
		}
		dfs(node.Left, curPath, curSum)
		dfs(node.Right, curPath, curSum)

		curPath = curPath[:len(curPath)-1]
	}

	dfs(root, []int{}, 0)
	return ans
}

// bfs 思路
// 1.使用map存储每个节点与其父节点的映射关系, 以及每个节点剩余需要扣减的值
// 2.当遇到子节点并且路径值合适时, 从map中倒推父节点链路
func pathSum113_bfs(root *TreeNode, target int) [][]int {
	ans := [][]int{}

	if root == nil {
		return ans
	}

	// 存储一个节点与其父节点的映射
	nodeParentMap := make(map[*TreeNode]*TreeNode)
	nodePathSumMap := make(map[*TreeNode]int)
	nodeParentMap[root] = nil
	nodePathSumMap[root] = root.Val

	var buildPath func(*TreeNode) []int
	buildPath = func(node *TreeNode) []int {
		path := []int{}
		for node != nil {
			path = append([]int{node.Val}, path...)
			node = nodeParentMap[node]
		}
		return path
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		if node.Left == nil && node.Right == nil && nodePathSumMap[node] == target {
			// find path
			ans = append(ans, buildPath(node))
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
			nodeParentMap[node.Left] = node
			nodePathSumMap[node.Left] = nodePathSumMap[node] + node.Left.Val
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			nodeParentMap[node.Right] = node
			nodePathSumMap[node.Right] = nodePathSumMap[node] + node.Right.Val
		}

		queue = queue[1:]
	}
	return ans
}
