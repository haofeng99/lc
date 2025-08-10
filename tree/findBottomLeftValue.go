package tree

// https://leetcode.cn/problems/find-bottom-left-tree-value/
// 513. 找树左下角的值
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return -1
	}

	maxDeep := -1
	leftNodeVal := -1

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, deep int) {
		if node == nil {
			return
		}
		if deep > maxDeep {
			leftNodeVal = node.Val
			maxDeep = deep
		}
		dfs(node.Left, deep+1)
		dfs(node.Right, deep+1)
	}

	dfs(root, 1)
	return leftNodeVal
}

func findBottomLeftValue_bfs(root *TreeNode) int {
	if root == nil {
		return -1
	}
	leftNodeVal := -1
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		nums := len(queue)
		first := true
		for nums > 0 {
			node := queue[0]
			if first {
				leftNodeVal = node.Val
				first = false
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
			nums--
		}
	}

	return leftNodeVal
}
