package tree

// https://leetcode.cn/problems/binary-tree-maximum-path-sum/
// 124. 二叉树中最大路径和
// 当前节点有两个角色
// 1.路径上的节点
// 2.路径上节点的子节点
// var maxSum int

// func maxPathSum(root *TreeNode) int {
// 	maxSum = -1001
// 	dfs_maxPathSum(root)
// 	return maxSum
// }

// func dfs_maxPathSum(node *TreeNode) int {
// 	if node == nil {
// 		return 0
// 	}
// 	leftSum := dfs_maxPathSum(node.Left)
// 	rightSum := dfs_maxPathSum(node.Right)

// 	// 1.路径上的节点
// 	currPathValSum := node.Val + leftSum + rightSum
// 	if currPathValSum > maxSum {
// 		maxSum = currPathValSum
// 	}

// 	// 2.路径上节点的子节点
// 	return max(node.Val+max(leftSum, rightSum), 0)
// }

// ----------------------------------

var maxSumPath_ []int
var maxSum_Path int

func maxPathSumPath(root *TreeNode) int {
	maxSum_Path = -1001
	_, _ = dfs_maxPathSumPath(root)
	return maxSum_Path
}

func dfs_maxPathSumPath(node *TreeNode) (int, []int) {
	if node == nil {
		return 0, []int{}
	}
	leftSum, leftPath := dfs_maxPathSumPath(node.Left)
	rightSum, rightPath := dfs_maxPathSumPath(node.Right)

	// 1.路径上的节点
	currPathValSum := node.Val + leftSum + rightSum
	if currPathValSum > maxSum_Path {
		maxSum_Path = currPathValSum

		path := []int{}
		if leftSum > 0 {
			path = append(path, leftPath...)
		}
		path = append(path, node.Val)
		if rightSum > 0 {
			path = append(path, rightPath...)
		}

		maxSumPath_ = path
	}

	// 2.路径上节点的子节点
	if leftSum > rightSum {
		return max(node.Val+leftSum, 0), append(leftPath, node.Val)
	} else {
		return max(node.Val+rightSum, 0), append([]int{node.Val}, rightPath...)
	}
}

func dfs_maxPathSum(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftSum := dfs_maxPathSum(node.Left)
	rightSum := dfs_maxPathSum(node.Right)

	// 1. 路径上的一个节点
	curPathSum := leftSum + rightSum + node.Val
	if curPathSum > maxSum_Path {
		maxSum_Path = curPathSum
	}

	// 2.路径上的子节点
	return max(node.Val+max(leftSum, rightSum), 0)
}
