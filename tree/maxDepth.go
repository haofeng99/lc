package tree

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/
// 104 二叉树的最大深度
var maxDepthVal int

func maxDepth(root *TreeNode) int {
	maxDepthVal = 0
	if root == nil {
		return maxDepthVal
	}

	dfs_maxDepth(root, 1)
	return maxDepthVal
}

func dfs_maxDepth(node *TreeNode, depth int) {
	if node == nil {
		return
	}
	if depth > maxDepthVal {
		maxDepthVal = depth
	}
	dfs_maxDepth(node.Left, depth+1)
	dfs_maxDepth(node.Right, depth+1)
}

func bfs_maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	nodeList := []*TreeNode{root}
	ans := 0
	for len(nodeList) != 0 {
		num := len(nodeList)
		for num > 0 {
			node := nodeList[0]
			if node.Left != nil {
				nodeList = append(nodeList, node.Left)
			}
			if node.Right != nil {
				nodeList = append(nodeList, node.Right)
			}
			nodeList = nodeList[1:]
			num--
		}
		ans++
	}

	return ans
}
