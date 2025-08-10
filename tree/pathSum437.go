package tree

// https://leetcode.cn/problems/path-sum-iii
// 437. 路径总和 III
// 对每个节点都做dfs , 前缀和
var path_III [][]int
var count_Path int

func pathSum437(root *TreeNode, targetSum int) int {
	path_III = [][]int{}
	count_Path = 0
	pathSum_III(root, targetSum)
	return count_Path
}
func pathSum_III(root *TreeNode, targetSum int) {
	if root == nil {
		return
	}
	dfs_pathSum_III(root, targetSum, []int{})

	pathSum_III(root.Left, targetSum)
	pathSum_III(root.Right, targetSum)
}

func dfs_pathSum_III(node *TreeNode, targetSum int, curPath []int) {
	if node == nil {
		return
	}

	curPathSum := targetSum - node.Val
	curPath = append(curPath, node.Val)
	if curPathSum == 0 {
		path_III = append(path_III, append([]int{}, curPath...))
		count_Path++
	}

	dfs_pathSum_III(node.Left, curPathSum, curPath)
	dfs_pathSum_III(node.Right, curPathSum, curPath)

	// 回溯
	curPath = curPath[:len(curPath)-1]
}
